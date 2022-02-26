module ContactForm exposing (main)

import Browser
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events
import Http
import Json.Encode


-- MAIN


main : Program String Model Msg
main =
    Browser.element
        { init = \_ -> ( Model (View form) "" "" "" "" "" "" "" "", Cmd.none )
        , update = update
        , subscriptions = \_ -> Sub.none
        , view = view
        }



-- MODEL


type alias Model =
    { view : View
    , fullName : String
    , phone : String
    , email : String
    , company : String
    , position : String
    , employeeCount : String
    , hear : String
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
    Html.form [ Html.Events.onSubmit Submit ]
        [ div [ class "row" ]
            [ div [ class "col-12" ]
                [ div [ class "form-group" ]
                    [ label [ for "fullName" ] [ text "Full name*" ]
                    , input [ class "form-control", id "fullName", placeholder "Enter your full name", type_ "text", required True, Html.Events.onInput (Input FieldFullName) ] []
                    ]
                ]
            , div [ class "col-12 col-md-6" ]
                [ div [ class "form-group" ]
                    [ label [ for "phone" ]
                        [ text "Phone*" ]
                    , input [ class "form-control", id "phone", placeholder "Enter your phone", type_ "text", required True, Html.Events.onInput (Input FieldPhone) ]
                        []
                    ]
                ]
            , div [ class "col-12 col-md-6" ]
                [ div [ class "form-group" ]
                    [ label [ for "email" ] [ text "Email*" ]
                    , input [ class "form-control", id "contactName", placeholder "Enter your email", type_ "email", required True, Html.Events.onInput (Input FieldEmail) ] []
                    ]
                ]
            , div [ class "col-12 col-md-6" ]
                [ div [ class "form-group" ]
                    [ label [ for "company" ] [ text "Company*" ]
                    , input [ class "form-control", id "contactName", placeholder "Enter your company", type_ "text", required True, Html.Events.onInput (Input FieldCompany) ] []
                    ]
                ]
            , div [ class "col-12 col-md-6" ]
                [ div [ class "form-group" ]
                    [ label [ for "contactName" ]
                        [ text "Position*" ]
                    , input [ class "form-control", id "contactName", placeholder "Enter your position", type_ "text", required True, Html.Events.onInput (Input FieldPosition) ]
                        []
                    ]
                ]
            , div [ class "col-12 col-md-6" ]
                [ div [ class "form-group" ]
                    [ label [ for "contactName" ] [ text "No. of employees" ]
                    , select [ class "form-control", Html.Events.onInput (Input FieldEmployeeCount) ]
                        [ option [ disabled True, selected True ] [ text "Select an option" ]
                        , option [] [ text "1-50 employees" ]
                        , option [] [ text "51-200 employees" ]
                        , option [] [ text "201-500 employees" ]
                        , option [] [ text "501-1000 employees" ]
                        , option [] [ text "1001-5000 employees" ]
                        , option [] [ text "5001-10,000 employees" ]
                        , option [] [ text "10,001+ employees" ]
                        ]
                    ]
                ]
            , div [ class "col-12 col-md-6" ]
                [ div [ class "form-group" ]
                    [ label [ for "contactName" ] [ text "How did you hear about us?" ]
                    , select [ class "form-control", Html.Events.onInput (Input FieldHear) ]
                        [ option [ disabled True, selected True ] [ text "Select an option" ]
                        , option [] [ text "Search Engine" ]
                        , option [] [ text "Google" ]
                        , option [] [ text "Facebook" ]
                        , option [] [ text "YouTube" ]
                        , option [] [ text "LinkedIn" ]
                        , option [] [ text "Twitter" ]
                        , option [] [ text "Instagram" ]
                        , option [] [ text "Other social media" ]
                        , option [] [ text "Email" ]
                        , option [] [ text "Word of mouth" ]
                        , option [] [ text "Other" ]
                        ]
                    ]
                ]
            , div [ class "col-12" ]
                [ div [ class "form-group" ]
                    [ label [ for "contactName" ] [ text "Message" ]
                    , textarea [ class "form-control", id "contactMessage", placeholder "Tell us what we can help you with!", attribute "rows" "5", Html.Events.onInput (Input FieldMessage) ] []
                    ]
                ]
            , div [ class "col-12" ]
                [ div [ class "form-group mb-0" ]
                    [ button [ class "btn btn-block btn-primary lift bg-secondary form-submit-btn" ]
                        [ text "Send message" ]
                    ]
                ]
            ]
        ]



-- UPDATE


type Msg
    = Submit
    | Input Field String
    | Submitted (Result Http.Error ())


type Field
    = FieldFullName
    | FieldEmail
    | FieldPhone
    | FieldCompany
    | FieldPosition
    | FieldEmployeeCount
    | FieldHear
    | FieldMessage


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        Input field s ->
            let
                modelUpdated =
                    case field of
                        FieldFullName ->
                            { model | fullName = s }

                        FieldEmail ->
                            { model | email = s }

                        FieldPhone ->
                            { model | phone = s }

                        FieldCompany ->
                            { model | company = s }

                        FieldPosition ->
                            { model | position = s }

                        FieldEmployeeCount ->
                            { model | employeeCount = s }

                        FieldHear ->
                            { model | hear = s }

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
            , Http.post
                { url = "/.netlify/functions/contact-us"
                , body =
                    Http.jsonBody
                        (Json.Encode.object
                            [ ( "Message"
                              , Json.Encode.string
                                    (String.join "\n"
                                        [ "Contact Form Submission"
                                        , ""
                                        , "Full Name: " ++ model.fullName
                                        , "Email: " ++ model.email
                                        , "Phone: " ++ model.phone
                                        , "Company: " ++ model.company
                                        , "Position: " ++ model.position
                                        , "Employee Count: " ++ model.employeeCount
                                        , "How did you hear about us?: " ++ model.hear
                                        , ""
                                        , "Message:"
                                        , ""
                                        , model.message
                                        ]
                                    )
                              )
                            ]
                        )
                , expect = Http.expectWhatever Submitted
                }
            )

        Submitted result ->
            case result of
                Ok () ->
                    ( { model | view = View submitted }, Cmd.none )

                Err _ ->
                    ( { model | view = View failed }, Cmd.none )
