module JobForm exposing (main)

import Base64.Encode
import Browser
import Bytes
import Bytes.Encode
import File
import File.Select
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Json.Encode
import Task



-- MAIN


main : Program String Model Msg
main =
    Browser.element
        { init = \_ -> ( Model (View form) Nothing "" "" "", Cmd.none )
        , update = update
        , subscriptions = \_ -> Sub.none
        , view = view
        }



-- MODEL


type alias Model =
    { view : View
    , maybeFile : Maybe File.File
    , fullName : String
    , email : String
    , message : String
    }


type View
    = View (Model -> Html.Html Msg)



-- VIEW


view : Model -> Html.Html Msg
view model =
    (\(View toElement) -> toElement model) model.view


submitted : Model -> Html.Html Msg
submitted model =
    h3 [ class "text-center" ] [ text "Thank you for submitting! We have received your message. Kindly wait for us to get back to you." ]


failed : Model -> Html.Html Msg
failed model =
    h3 [ class "text-center" ] [ text "Whoops! An error occurred. Please refresh the page and fill up the form again. Sorry for the inconvenience." ]


form : Model -> Html.Html Msg
form model =
    div [ class "card card-border bmrder-primary shadow-light-lg" ]
        [ div
            [ class "card-body" ]
            [ Html.form [ Html.Events.onSubmit Submit ]
                [ div [ class "row" ]
                    [ div [ class "col-12 col-md-6" ]
                        [ div [ class "form-group mb-5" ]
                            [ label [ for "applyFullName" ] [ text "Full name" ]
                            , input [ type_ "text", class "form-control", required True, id "applyFullName", placeholder "Full name", onInput (Input FieldFullName) ] []
                            ]
                        ]
                    , div [ class "col-12 col-md-6" ]
                        [ div [ class "form-group mb-5" ]
                            [ label [ for "applyEmail" ] [ text "Email" ]
                            , input [ type_ "email", class "form-control", required True, id "applyEmail", placeholder "Enter your email", onInput (Input FieldEmail) ] []
                            ]
                        ]
                    , div [ class "col-12" ]
                        [ div [ class "form-group mb-5" ]
                            [ p [ class "mb-2" ] [ text "Attachment (must be less than 5mb and file type is JPEG/PNG/PDF)" ]
                            , a [ class "btn btn-primary", href "#file", onClick SelectFile ]
                                [ text
                                    (case model.maybeFile of
                                        Just file ->
                                            File.name file

                                        Nothing ->
                                            "Choose file"
                                    )
                                ]
                            ]
                        ]
                    ]
                , div [ class "row" ]
                    [ div [ class "col-12" ]
                        [ div [ class "form-group mb-5" ]
                            [ label [ for "applyMessage" ] [ text "Message" ]
                            , textarea [ id "applyMessage", rows 5, class "form-control", placeholder "(Optional)", onInput (Input FieldMessage) ] []
                            ]
                        ]
                    ]
                , div [ class "row align-items-center" ]
                    [ div [ class "col-12 col-md" ]
                        [ button [ class "btn btn-primary mb-6 mb-md-0 lift" ]
                            [ text "Submit", i [ class "fe fe-arrow-right ml-3" ] [] ]
                        ]
                    , div [ class "col-12 col-md-auto" ]
                        [ p [ class "font-size-sm text-muted mb-0" ]
                            [ text "Your application will be sent securely and remain private." ]
                        ]
                    ]
                ]
            ]
        ]



-- UPDATE


type Msg
    = Submit
    | SelectFile
    | FileSelected File.File
    | Input Field String
    | Submitted (Result Http.Error ())


type Field
    = FieldFullName
    | FieldEmail
    | FieldMessage


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        SelectFile ->
            ( model, File.Select.file [ "image/png", "image/jpeg", "application/pdf" ] FileSelected )

        FileSelected file ->
            ( { model | maybeFile = Just file }, Cmd.none )

        Input field s ->
            let
                modelUpdated =
                    case field of
                        FieldFullName ->
                            { model | fullName = s }

                        FieldEmail ->
                            { model | email = s }

                        FieldMessage ->
                            { model | message = s }
            in
            ( modelUpdated, Cmd.none )

        Submit ->
            let
                _ =
                    Debug.log "Debug.log" "submitted!"
            in
            ( model
            , Task.attempt Submitted
                (Task.andThen
                    (\attachments ->
                        Http.task
                            { method = "POST"
                            , headers = []
                            , url = "/.netlify/functions/contact-us"
                            , body =
                                Http.jsonBody
                                    (Json.Encode.object
                                        [ ( "Message"
                                          , Json.Encode.string
                                                (String.join "\n"
                                                    [ "Job Application for Account Sales Executive"
                                                    , ""
                                                    , "Full Name: " ++ model.fullName
                                                    , "Email: " ++ model.email
                                                    , ""
                                                    , "Message:"
                                                    , ""
                                                    , model.message
                                                    ]
                                                )
                                          )
                                        , ( "Attachments", attachments )
                                        ]
                                    )
                            , resolver = Http.stringResolver (\_ -> Ok ())
                            , timeout = Nothing
                            }
                    )
                    (case model.maybeFile of
                        Just file ->
                            Task.andThen
                                (\bytes ->
                                    Task.succeed
                                        (Json.Encode.list
                                            (\f ->
                                                Json.Encode.object
                                                    [ ( "Name", Json.Encode.string (File.name f) )
                                                    , ( "Content", Json.Encode.string (Base64.Encode.encode (Base64.Encode.bytes bytes)) )
                                                    , ( "ContentType", Json.Encode.string (File.mime f) )
                                                    ]
                                            )
                                            [ file ]
                                        )
                                )
                                (File.toBytes file)

                        Nothing ->
                            Task.succeed (Json.Encode.list (\_ -> Json.Encode.null) [])
                    )
                )
            )

        Submitted result ->
            case result of
                Ok () ->
                    ( { model | view = View submitted }, Cmd.none )

                Err _ ->
                    ( { model | view = View failed }, Cmd.none )
