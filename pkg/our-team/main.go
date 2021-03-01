package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	t, err := template.ParseFiles("template.tmpl")
	if err != nil {
		panic(err)
	}

	if err := t.ExecuteTemplate(os.Stdout, os.Args[1], map[string]interface{}{
		"Members": members(),
	}); err != nil {
		panic(err)
	}
}

type member struct {
	LastName  string
	FirstName string
	Position  string
	Bio       []string
	Slug      string
}

func members() []member {
	ms := []member{}

	for _, m := range [][]string{
		{"Syquia", "Nenuca", "CEO, Managing Partner", "Nenuca is an organizational effectiveness practitioner who specializes in organization design, team coaching, and executive facilitation. She co-founded Better Organizations by Design with the mission of enabling small and medium businesses achieve scale, disrupt industries, and create currently unimagined value for people. \n She has driven change through various HR roles in the tech and professional service industries, as well at a large multi-industry conglomerate. Among many projects, Nenuca’s favorites include redesigning engineering and services organizations as part of business transformations and helping companies reshape their culture with intention. \n She has a Master’s in Industrial and Labor Relations from Cornell and a Master of Science in Organization Development from the University of San Francisco. She is also a certified design thinking instructor with LUMA Institute.  "},
		{"Petrovic", "Lydia", "Co-founder, Partner", "Lydia is a team effectiveness specialist with deep roots in the nonprofit world. She co-founded Better Organizations by Design driven by a desire to enable every person at every level of the organization to be their best selves at work. \n Lydia’s specialties are team interventions, talent development with a diversity, equity and inclusion lens, and lateral process design. Her experience in sales development and nonprofit management have made her adept at creating opportunity and solving wicked problems in resource-constrained environments.  \n Having lived and worked in China and Russia, she brings a deep appreciation of wide-ranging perspectives to every discussion. With a background in linguistics and expertise in translation and localization, she also guides clients on the role of language in shaping organizational culture. \n She has a Masters in Organization Development from the University of San Francisco, as well as a degree in Russian and Linguistics from UC Berkeley.\n As a new mom of twins (and a change management expert), Lydia has a renewed purpose to help others navigate the joys and pains of big transitions."},
		{"Wingo", "Brett", "Business Advisor, Partner", "Brett is a 30-year veteran of notable Silicon Valley technology companies. He started his career at Apple, then co-founded a start-up that was acquired after a successful product launch. He enjoyed the start-up world for many years as a senior executive and finished with two acquisition exits.\n In 2008, Brett joined Cisco and held many senior leadership positions, including the GM and CEO role for Linksys.  Following the turnaround and sale of Linksys, he served as GM for several business groups, with full product and P&L responsibility of businesses from $500M to $4.6B in annual revenues. He also led the transformation of Cisco’s $12B services business into a recurring revenue model built on cutting-edge customer success methodologies and modern software.    \nKnown as a transformational leader with a keen sense for how to accelerate the future potential of a business, he has increased profit and growth substantially in every business he has led.  His passion for business strategy, building great teams, and developing leaders have led him into his new career in consulting, teaching and business advisory.  "},
		{"Abdullah", "Zahra", "Consultant", "Zahra has over 17 years of global experience as a change management advisor at top firms, including McKinsey & Co.\nShe has worked in diverse countries, including Saudi Arabia, Dubai, the USA, and Canada, adeptly navigating complex cultural differences and frequently shifting professional demands. She has successfully led national transformation programs and initiatives across various sectors such as health, education, financial, government services, and non-profit organization.\nZahra's true passion is coaching and leadership development, which she discovered during her work in the field of change management. Her life's work is to help leaders and their teams embrace change as an opportunity for growth. She brings a depth of insight into managing multicultural task forces, creating an inclusive culture, fostering leadership effectiveness, and facilitating successful business transformations.\nShe is an ICF certified executive coach and holds an Executive MBA from Hult International Business School."},
		{"Amoukhteh", "Peiman", "Consultant", "Peiman is a data scientist with over 20 years of experience using predictive modeling, data mining, simulation development and data visualization to solve challenging business problems. He has worked with many leading high technology companies such as HPE, Microsoft, NetApp, VMware, Seagate as well as others. \n He started his career as a reliability and quality engineer for Solectron Corporation (now Flextronics) and became the head of their Advanced Technology Group.  Subsequently, he founded S2I2, a software company for quality management.  He also worked for Qtron Corporation and Comtel Electronics as their VP of Operations and COO, respectively.  In 1995 he received the Entrepreneur of the Year award from SBA in Northern California.\n Peiman holds both a Bachelor’s and Master’s degree in Physics.  "},
		{"Augustin", "Sally", "Consultant", "Sally Augustin, PhD, is a practicing environmental/design psychologist. She has extensive experience integrating science-based insights to develop recommendations for the design of places, objects, and services that support desired cognitive, emotional, and physical experiences.  \n Her clients include manufacturers, service providers, and design firms in North America, Europe, and Asia.  \n Dr. Augustin is also a Fellow of the American Psychological Association."},
		{"Bland", "Ken", "Project Manager", "Ken has spent over three decades as a program manager and team leader in the software industry, leading and designing small strategic executive events to large multi-million dollar conferences. He has served clients from a broad cross-section of industries, domestic and international. His core expertise lies in taking concepts and building plans to convert vision into reality.\n Ken’s passion is focused on giving back and making a difference by actively holding several leadership roles with youth-based organizations focus on leadership development and historically underserved children. Ken also supports and volunteers at a local animal shelter.  "},
		{"Burns", "Collene", "Consultant", "Collene is an experienced facilitator who is passionate about helping organizations identify and achieve their big, hairy, audacious goals. \n While working in the high tech, healthcare, not-for-profit, manufacturing, higher education and services sectors, Collene developed techniques that spark innovation and promote collaboration in focused, facilitated settings.\n At companies such as Microsoft and Aramark, and at various not-for-profits, Collene has brought a focus on solving business problems through the human elements of change management in Human Resources and Operations.\n Collene supports high potential people and organizations through leadership coaching and facilitation. She also provides conflict resolution services with coaching, training, and mediation for workplaces and families.\n She holds a B.S. in Industrial and Labor Relations from Cornell University and an MBA from Simon Business School (University of Rochester).\n Collene calls Rochester, N.Y. home base where she and her husband spend a fair amount of time cheering on their kiddo at Nordic skiing and golf competitions. They consider Seattle, WA their second hometown, where Collene began her search for the elusive perfect cup of coffee.  "},
		{"Burns", "Zachary", "Consultant", "Zach works at the intersection of psychology and business. Leveraging his expertise on patterns of human cognition and how they can bridge the gap between ideal and actual behavior, he helps drive transformations across our client organizations.\n He currently teaches at University of San Francisco School of Management and previously taught at Kellogg School of Management. His courses include Managerial Decision Making, Negotiations, Building and Leading Effective Teams, and Management and Organizational Dynamics. \n His research has been published in Proceedings of the National Academy of Sciences and Journal of Experimental Psychology: General, and has appeared in the New York Times, Hidden Brain, Slate, and MTV Finland. \n He received his MBA and PhD from Booth School of Business in Managerial and Organizational Behavior with a focus in Economics.  "},
		{"Bruce", "Veronica", "Director of Finance & Operations", "Veronica is a future-focused financial leader with over 25 years of experience honing her advisory skills with Fortune 500 companies like PwC and Sprint. Her diverse industry experience spans telecom, non-profit, publishing, property management, construction, government, and the technology sector.\nShe has a keen eye for detail and an unparalleled ability to develop processes that increase organizational effectiveness so that individuals and businesses can thrive together.\nVeronica graduated with a Bachelor of Science degree from Missouri State University. She lives in gulf coastal Alabama with her husband Alfred and enjoys sharing in bible ministry and spending time with her children, family and friends."},
		{"Cohen", "Alison", "Consultant", "As a researcher and statistician with experience in large scale communities and a background outside of corporate HR, Alison gives our clients a very fresh take on their data. We leverage her deep expertise to ensure that quantitative and qualitative analyses are designed to produce quality insights.\n She is also extremely passionate about the social justice implications of big data. \n Alison received her Masters in Public Health in Epidemiology and Biostatistics and her PhD in Epidemiology from University of California, Berkeley.  "},
		{"Coppock", "Abby", "Consultant", "Abby helps leaders build creative, inclusive organizations where everyone can thrive. With over 17 years of experience across industries, Abby specializes in organizational transformation, leadership development, and change management.\nAbby holds equity in the center and is mindful of how processes, incentives and everyday behavior produce organizational outcomes. She is an engaging facilitator and strategic thinker who helps organizations clarify their strategy and values and put them into action.  Empathy, curiosity, and participatory design are her superpowers.\nAbby has worked with clients including Nike, the Bill and Melinda Gates Foundation, the Port of Portland, and Bonneville Power Administration. Prior to consulting, Abby led internal and external communications for the City of Portland, Office of Management and Finance.\nStarting her career as a social worker, Abby’s care for people is evident in the way she listens, reads non-verbal cues, and is inclusive of perspectives that are not always considered. She has an MA in Social Service Administration from the University of Chicago, with bachelor degrees in Communications and Economics from Wheaton College. In her spare time, Abby loves to paint and stays involved with local government initiatives."},
		{"de Leon", "Maria", "", "Maria has led over a hundred projects to successful delivery with over 15 years of experience working with insurance companies, tax collectors, utilities, communications and non-profit organizations.  Experienced in all aspects of software development and roll-out, Maria is technical enough to debug code, but still relatable when collaborating with end-users.  While she is an adept subject-matter-expert, she also knows what it takes to rapidly build the kind of trust required to collaborate effectively with clients and team members alike.\nA ‘MacGyver’ of sorts, Maria provides effective solutions leveraging scarce resources, then builds on past experiences to ensure that lessons are documented and used to better future performance.  Maria especially enjoys resolving hard-to-resolve issues with creativity and an eye for both short and long-term considerations.\nMaria is fulfilled by lifting people up, especially in their times of greatest need. In her spare time, she volunteers at her children's school to expand recycling efforts and reduce waste."},
		{"Delgado", "Mark", "Designer", "Mark is a visual design expert with nearly 10 years of experience in graphic design, art direction and gamification. He has extensive experience consulting for the video game industry, planning and executing strategic design initiatives to elevate marketing and communications campaigns.\nMark has a degree in Computer Science with a specialization in Game Design and Development. He loves the art of visual expression, and is always looking for outlets to flex his animation chops. A native of Manila, Mark currently resides in Caloocan City. In his spare time, he loves taking care of his growing family and taking opportunities to be the best at whatever new activity he tries."},
		{"Dyer", "Tina", "Consultant", "Tina is a financial services professional with over three decades of experience leading in areas of global racial equity, diversity, and inclusion, public speaking. As a seasoned learning and development leader, she has been equipping leaders to engage in crucial conversations and difficult subjects for more than a decade. Tina specializes in creating brave spaces where people can feel psychologically safe to engage, explore and take action to achieve the outcomes important to them.\nPreviously, Tina served as Vice President and DE&I Learning Strategy Consultant for a Fortune 50 financial services company, where she designed and implemented data-driven enterprise-wide strategies and programs to increase employee engagement and productivity.\nTina has a B.S. in Psychology from University of California at Davis, and a Master’s degree in Spiritual Formation and Leadership from Friends University. She is a Courageous Conversation about Race practitioner and in training as a Somatic Abolitionist. Tina is also certified to coach to the Intercultural Development Inventory (IDI) and Inclusion Skills Measurement (ISM) assessments, both designed to raise awareness and support intercultural competence development.\nTina treasures being Mom to two young-adult children and Auntie to 62 nieces and nephews! She is a firm believer that exposure to the arts is an essential part of healthy development."},
		{"Fiorentino", "Dina", "Project Manager", "Dina Fiorentino spent 30 years in the technology industry working for companies such as SAS Institute, GTE Wireless and Apple Computer. She has led large teams, built marketing organizations, focused on team development, nurtured a passion for project management and pursued her business coaching certification. Dina also serves as a volunteer mentor/coach to small business and nonprofit leaders as well as young professionals."},
		{"Frederick", "LT", "Consultant", "LaTricia is a senior HR professional and certified leadership coach with over 20 years of experience helping individuals, teams and leaders achieve optimal performance.\n She has extensive experience in operations, talent management and strategic planning in the worlds of high-tech and education. A lover of learning, she has also been an adjunct professor for management studies.\n Since 2015, she has dedicated herself to driving organizational improvements with a diversity and inclusion lens.\n She has an MBA with a focus in HR and is currently working towards a Doctor of Education in Leadership and Learning in Organizations.  "},
		{"Frazer", "Eric", "Consultant", "Eric is a psychologist with 20 years (and over 10,000 hours) of expertise in psychological assessment. His passion is for uncovering the behaviors and personality traits that lead to successful personal and professional outcomes.\nHe has a deep understanding of the psychological factors that play a role in high performing leaders who are loved by their teams. Eric has provided decision-makers with proven techniques and tools to identify key characteristics when recruiting top talent. He also has a proven record of accelerating professional development for high performers and top leaders using actionable coaching and training methods.\nEric has spent many years carefully curating psychological research, writing about its applications on human performance, and developing strategies and protocols to help top professionals consistently push the boundaries of their human potential. In his spare time, he pursues physical and psychological challenges in the \"outdoor lab\" in order to enhance his own performance and productivity."},
		{"Godfrey", "Alison", "", "Alison Godfrey is an Executive Coach with an impressive list of international clients. Much of her time is spent with start-up CEOs who are new to the world of entrepreneurship, helping those leaders navigate the balance of the internal and external facing aspects of an organization.\nOver the past 45 years, Alison has been both an entrepreneur and an experienced corporate C-suite executive. She has founded and led international start-ups in fields ranging from health, metallurgy, to condensed matter physics. As CEO for Concentric Advisors, TX she developed and led a team of top physicists, electrochemists, and material scientists to create the premier research laboratory in the global field. She also founded and was CEO for Energetics Technologies in Israel, where she established highly successful cooperative research projects with Italy’s nuclear research laboratory (ENEA) as well as SRI (Stanford Research Institute) in CA, funded by grants from DARPA (U.S. Defense Projects Research Agency). Alison has worked closely with some of the world’s most influential individuals in technology, business leadership, and philanthropy.\nAs a dynamic, highly effective, and goal-oriented change agent with a diverse background, she has enhanced the growth and development of countless individuals and organizations. Alison is recognized as a highly effective communicator capable of cultivating partnerships across all lines of business to promote cohesive business practice and produce high-impact outcomes.\nHer vast corporate experience has culminated in a flourishing career in executive coaching and innovative leadership development. Alison’s warmth and personal approach facilitates individuals to shift their thinking, challenge paradigms, remove blocks, and ultimately change behaviors. This allows her clients to enhance their self-awareness and self confidence, achieve success, and make a positive impact on their communities.\nAlison retains a Certificate of Business Excellence from the Haas School of Business (UC Berkeley) and an Executive Coach certification from the Berkeley Coaching Institute. She is also an Assistant Lecturer at the Haas School of Business, where she teaches Executive Leadership and Communication to Executive MBA candidates."},
		{"Green", "Jennifer", "Consultant", "Jennifer is a rewards expert with decades of global experience. Her career has covered small to mid-size public and pre-public companies to global Fortune 50 companies in high-tech, life science, and health insurance industries.\n Beyond designing reward systems to drive behavior and performance, she also excels in leading clients through the change process, from preparation to sustainable implementation.  "},
		{"Hunter", "Keith", "Consultant", "Keith is a professor of organizational behavior as well as a management consultant, with a passion for the application of behavioral and system science to both superior organizational outcomes and human well-being.\nKeith operates at the intersection of cognition, structure and social dynamics. His work most commonly tackles the challenges of developing leadership, transforming culture, and achieving team effectiveness. His consulting work has produced a leadership case, leader training in DEIB and allyship, and workshops on commitment and change. Keith’s network analyses have also supported gender equity initiatives and organizational restructuring.\nHe currently researches network perceptions and situational constraints, which are critical factors to organizational progress. Keith has earned both a Ph.D. in Organizational Behavior and Management and a Master of Philosophy in Public Policy at Carnegie Mellon University. He has also served in the US Navy, completed Bachelors and Masters degrees in computer science, and worked a software engineer within national laboratory and start-up environments. In his free time he enjoys going on low-key nature hikes, playing tennis, and exploring food and travel adventures."},
		{"Jenkins", "Jenine Smith", "Consultant", "Jenine is an award-winning consultant with broad experience in HR and a special passion for leadership development and diversity, equity and inclusion. She has worked with family-owned and Fortune 100 companies, always with the mission of connecting women from diverse backgrounds with opportunity and a sense of belonging. She has also served as interim HR leader for growing organizations. \n Jenine is working towards her Professional Certified Coach accreditation with the International Coaching Federation (ICF).  "},
		{"Jessen", "Martin", "Consultant", "Martin is an international HR leader with a track record of aligning business transformations, developing leaders, building successful teams and shaping culture.\n His expertise includes organization design, strategic workforce planning, people analytics, talent development, and performance management. Martin’s roles as an HR Business Partner and a functional leader in companies like IBM and Schneider Electric have taken him to Germany, Denmark, England, France, and the United States. \n He is passionate about enabling HR teams to be increasingly strategic and able to drive innovative and impactful solutions globally. \n He has an MBA with a concentration in HR."},
		{"Jordan", "Marty", "Consultant", "Marty is an accomplished and versatile senior organization development practitioner with broad-based experience in a variety of industries, working in diverse business functions, government and not-for-profit environments. \n She has worked at all organizational levels and is as comfortable working with a C-level executive as she is with a first-line supervisor or frontline employee. \n Her passion is helping leaders create productive workplaces where employees are challenged, engaged and bring their best selves to work every day. She has a BS in Education from Washington State University and a MS in OD from Pepperdine University."},
		{"Kapadia", "Danoosh", "", "Danoosh is a seasoned growth consultant who deeply understands the complex changes that leaders of organizations face. He works collaboratively with clients to transform their business and accelerate results.\nPrior to joining BOxD, Danoosh was VP of Business Development at XPLANE, and before that the Head of Business Development and Partnerships for IDEO U (the educational arm of the award-winning design and innovation design firm).\nDanoosh knows firsthand about the complex challenges organizations face as they grow and scale, since he helped build the company General Assembly—an education-technology start-up that teaches the most relevant and in-demand digital skills across data, design, business, and technology.\nBy training, Danoosh is an Industrial Engineer with experience designing industrial products for critical applications in oil and gas, chemical, and power where safety and reliability are of paramount importance. Danoosh holds a degree in Industrial and Systems Engineering with a minor in Business Administration from Virginia Tech. When he’s not working, you can find Danoosh at his home in San Francisco, building lego forts with his son and making chewy, blustery pizza."},
		{"Kromer", "Kate", "Consultant", "Kate designs and develops learning and development programs that communicate key business initiatives to employees to drive improved results. Her experience has focused on building engaging, interactive eLearning programs and training materials for businesses with complex global operations and distributed workforces.\n Kate collaborates with project stakeholders and subject matter experts to create content outlines, storyboards, voiceover scripts and course assessments. She develops media drafts in Articulate 360 tools that incorporate client-branded visuals and graphic elements, professional voiceover, conceptual graphics, custom animations, and learning interactions.  "},
		{"Leung", "Mark", "Consultant", "Mark is a corporate innovation practitioner who specializes in business design, strategy and leadership development. As the former Director of DesignWorks, the Business Design Lab at the Rotman School of Management he has trained thousands of business leaders how to innovate by design. \n He works with organizations to develop their internal innovation capabilities and catalyze new growth opportunities. Mark brings his engineering, business and human centred design expertise and love for making sense of wicked problems. He is passionate about empowering individuals with innovation agency and helping organizations transform to win. \n He has an MBA with a focus on Entrepreneurship and Strategy.  "},
		{"McCabe", "Caroline", "Consultant", "Caroline is a business psychologist who builds connections with clients to unleash the capabilities of an individual or a company’s workforce.\n Caroline has worked globally with a broad range of clients in technology, energy, healthcare, financial services, wealth management, consumer products and non profits. Caroline has held global roles internally and as a consultant in organization development, executive talent development and assessment and learning.\n She has a doctoral degree in psychology from the William James College of Psychology and completed her postdoctoral fellowship and internships via Harvard University Medical School. As a psychologist she uses the Hogan assessment instrument and is a Neurozone certified coach. \n At play, Caroline can be found adventuring with her family.  "},
		{"Nakar", "Grace", "Consultant", "Grace is a recognized HR leader who loves collaborative innovation. She has led change initiatives to elevate organization-wide talent development in hi-tech and financial services businesses, in the U.S. and Asia. She achieved this over and over again with a “bridgehead strategy” -- introducing systemic change in a target area where transformation gained a secure foothold, before spreading more widely. \nSince 2013, she has led enterprise-level Diversity and Inclusion & Learning, leveraging cross-functional teams to build DEI curricula and facilitator communities. She is passionate about working alongside organizations who want to match their intentions for DEI with their actions and impact. She provides holistic services with a consistent DEI lens - analytics, strategy, training, workplace reinforcement, talent management, and process change.\n She has an M.A. in Organizational Development & Leadership, with research on the intersection of DEI and adult stages of meaning making.  "},
		{"Neault", "Don", "Consultant", "As a seasoned executive coach and consultant, Don partners with his clients to enhance their awareness of personal and professional goals.  His respectful candor and his experience with globally diverse talent enables Don to help talented professionals visualize their future and plan for growth.  He is especially gifted in assisting in the development of early-in-career individuals, women and other diverse hi-potential talent who wish to maximize their career progression and navigate corporate environments.  Don has been recognized as a leader throughout the industry for developing women and under-represented leaders, and assisting them in dealing with the unique situations they face in the workplace.\nDon’s skills were developed and enhanced over a 40-year career spanning a variety of Fortune 100 firms. From traditional IT individual contributor roles, leadership positions, consulting, outsourcing business development, operations, to executive leadership positions, Don has truly worked a full spectrum of roles.  He has led teams small and large, and is known for developing strong teams and leaders with a variety of technical, business and organizational skills.\nMost recently, Don held multiple Vice President roles at Cisco Systems for 18 years, leading the management and delivery of consulting, engineering and professional services to diverse market segments-- including Financial Services, Manufacturing, Automotive, Logistics, Retail, Oil and Gas, Federal, State and Service Provider.  His clients and teams have spanned all regions of the globe.  In each case, he has driven a culture of focus on client satisfaction along with achieving revenue and profitability growth.  Don is known for his collaborative and inquisitive style, calm demeanor, and focus on results.\nDon has a BS in Mathematics from Lowell Technological Institute, an MS in Computer and Information Systems from Northeastern University, is a Certified Master Coach from the Center for Coaching Certification, as well as an ACC Credentialed Coach by the International Coaching Federation. He lives with his lovely wife of 44 years in Tryon, NC."},
		{"O'Reilly", "Meag-gan", "Consultant", "Meag-gan is a Staff Psychologist at Stanford University's Counseling and Psychological Services (CAPS) and Adjunct Faculty in the Stanford School of Medicine. As a former dancer, Dr. O’Reilly knows the power of the expressive arts and often weaves movement, writing and art therapy into her clinical work. While completing her Post-Doctoral Fellowship at CAPS, Dr. O'Reilly created the first satellite clinic for Black undergraduate and graduate students across the African Diaspora. She currently serves a Program Coordinator for Outreach, Equity, and Inclusion. In this role, Dr. O’Reilly co-created the Outreach and Social Justice Seminar in 2016 with the goal of training the next generation of culturally conscious and justice- oriented clinicians.\n In her work as a consultant, Dr. O'Reilly provides DEI consulting, healing spaces, workshops, and speaking engagements. Her experience spans non-profits as well as companies like Google, LYRA Health, and The United Negro College Fund's STEM Scholar Program that supports Black college students nationwide to navigate underrepresentation and discrimination in STEM fields.\n She also serves as the lead clinician in a partnership with Google to provide therapeutic spaces called. The Gathering Space for Black Google Employees in response to the murder of George Floyd and the chronic trauma, and grief, in the Black community. Her TEDx talk: Enough is Enough: The Power of Your Inherent Value, can be seen on YouTube and is a helpful reminder of unconditional self-worth and that our lives matter to the world."},
		{"Reichenburg", "Victoria", "Consultant", "For over 20 years, Victoria has been helping leaders author their own authentic stories, first as a filmmaker and now as an executive coach in the SF Bay Area.\nAs a past leader of creative teams for companies like Visa, Microsoft, Gap and Kaiser Permanente, she has mastered the challenge of prioritizing initiatives amidst resource limitations and competing demands. Along with being an ICF-certified coach, Victoria holds an MS in Organization Development and is trained in numerous tools including The Leadership Circle, Enneagram iEQ9, DISC, FIRO-B, and Immunity to Change.  Victoria’s methodology is grounded in a systems-level, evidence-based approach that delivers practical outcomes while honoring the human connection points that equip teams to thrive.\nVictoria lives in Oakland, CA with her active family, including her husband, two teenagers, their rescue dog, and two lazy cats."},
		{"Reogo", "Glaiza", "Executive Assistant", "As the CEO’s right hand, Glaiza wears many hats: project manager, administrative whiz, and a general tamer of moving parts.\n Beyond ensuring everything is well organized at BOxD, she adds value to the team by sharing her rich background as a licensed chemist for a cosmetic manufacturing company as well as laboratory management. \n Her mission is to provide virtual administrative and marketing assistance to coaches and consultants who advocate organizational development, authentic leadership, and managing people with a heart.  "},
		{"Stone", "Jade", "Consultant", "Jade is an organizational effectiveness specialist who instinctively blends data, a human-centered empathetic lens, organizational behavior, and systems-thinking. She brings a multidimensional perspective having worked across organizational settings including technology, startups, education, research, government, nonprofits, biotech, and international development. \n From her experience as a Peace Corps Volunteer in the South Pacific, to her time at IDEO developing digital tools that foster innovation and collaboration, Jade believes the power of effective teams can accomplish more than the sum of their parts.\n Her specialties include organizational alignment, strategy, designing an empowered environment, leveraging organizational strengths to inspire and drive change, and surfacing nuanced insights from data and interviews.\n Her on-the-ground experience is bolstered with a Masters in Organization Development (MSOD) from the University of San Francisco, where she graduated with honors.\n When she’s away from her desk, you can find her on the dance floor evangelizing her love of Lindy Hop (swing dancing), meditating, practicing Nonviolent Communication (NVC), and partnering in creative collaborations.  "},
		{"Thakkar", "Sejal", "Consultant", "As an employment lawyer with a teacher’s heart, Sejal ensures the solutions we co-create with clients are compliant, while increasing clients’ understanding of the ins and outs of the law. \n Her experience in employment-related matters includes an emphasis on wrongful termination, sexual harassment, discrimination, and retaliation. She counsels leaders, HR, and employees on all parts of the employment lifecycle.  "},
		{"Wells", "Audia", "Consultant", "Audia Wells is a specialist in change management,  HR and DEI solutions who has led critical initiatives building HR functions, facilitating mergers and acquisitions, transforming performance management systems, and implementing D&I strategies.\n Audia has supported a range of clients, from Fortune 500s to start-ups in nearly every industry: Entertainment, Professional Services, Manufacturing, Finance, Faith-based, Non-Profit, Pharma and  Healthcare, Higher Education, and Architecture/Design. \n As an engaging workshop and conference speaker/facilitator, Wells has shared her career and entrepreneurship expertise with Emory Univ Alumni Association, The Urban League of Greater Atlanta,  The Junior Leagues of Spokane, and Kane & DuPage Counties, Ladies Get Paid, and GECDC & Build  Bronzeville-IL SBDCs. \n Audia holds two master's degrees, an MA in International Development and an MA in Organizational Behavior. She is a DDI Certified facilitator, member of The Junior League, Delta Sigma Theta Sorority, Inc., Chicago SHRM  (Diversity, Inclusion and Belonging committee), Diversity and Inclusion Committee (Church of Jesus  Christ of Latter Day Saints), Illinois Diversity Council, Oxford College Alumni Board of Emory University, and is a member of the Association of Junior Leagues International DEI Certified Partners Workgroup."},
		{"Yusay", "April Love", "Administrative Assistant", "April is a seasoned operations and administration expert with over a decade of experience in the field of business process outsourcing. \n April has dedicated her time and talent to supporting business owners to develop and streamline the internal processes they need to be successful. She has provided customizable and scalable solutions to a wide range of clients, from CEOs at large companies to small business owners.\n A true renaissance woman, April is passionate about sustainable travel and tourism and also organizes gap year experiences for international travelers to the Philippines. From business development, marketing, to handling the nitty gritty of ground operations, she has truly done and seen it all. April loves to explore different places, cultures, and traditions, and enjoys connecting with international communities to exchange ideas and learning."},
	} {
		ms = append(ms, member{
			m[0],
			m[1],
			m[2],
			strings.Split(m[3], "\n"),
			slug(m[0], m[1]),
		})
	}

	return ms
}

func slug(ss ...string) string {
	for i, s := range ss {
		ss[i] = strings.TrimSpace(s)
	}

	s := strings.Join(ss, " ")
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ToLower(s)

	return s
}
