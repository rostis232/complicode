package handler

const (
	ua = "ua"
	en = "en"
)

var texts map[string]map[string]string
var team []TeamMember

type TeamMember struct{
	Photo string
	Name string
	Position string
	TwitteLink string
	FacebookLink string
	LinkedInLink string
}

func init() {
	texts = map[string]map[string]string{
		"title": {
			en: "CompliCode - Software Development Company | we solve your problems using code",
			ua: "CompliCode - розробка програмного забезпечення | ми вирішуємо ваші проблеми за допомогою коду",
		},
		"masthead": {
			ua: "Вітаємо у CompliCode!",
			en: "Welcome to CompliCode!",
		},
		"services": {
			ua: "Послуги",
			en: "Services",
		},
		"projects": {
			ua: "Проекти",
			en: "Projects",
		},
		"projectsDescription": {
			ua: "Наші проект",
			en: "Projects of our team.",
		},
		"contact": {
			ua: "Зв'язок",
			en: "Contact",
		},
		"team": {
			ua: "Команда",
			en: "Team",
		},
		"about": {
			ua: "Про нас",
			en: "About",
		},
		"client": {
			ua: "Клієнт",
			en: "Client",
		},
		"category": {
			ua: "Категорія",
			en: "Category",
		},
		"backend": {
			ua: "Бекенд",
			en: "Backend",
		},
		"frontend": {
			ua: "Фронтенд",
			en: "Frontend",
		},
		"db": {
			ua: "Бази даних",
			en: "Databases",
		},
		"deploy": {
			ua: "Деплой",
			en: "Deploy",
		},
		"close": {
			ua: "Закрити",
			en: "Close",
		},
		"webDevTitle": {
			ua: "Веб розробка",
			en: "Web Development",
		},
		"webDevDescription": {
			ua: "Ми надаємо комплексні послуги з розробки веб-додатків: від збору вимог зацікавлених сторін до розгортання та підтримки.",
			en: "We provide end-to-end web application development services from gathering stakeholder requirements to deployment and support.",
		},
		"parsingTitle": {
			ua: "Парсинг даних та веб-скрепінг",
			en: "Data Parsing & Web Scraping",
		},
		"parsingDescription": {
			ua: "Ми спеціалізуємося на вилученні та обробці даних з онлайн-джерел, перетворенні необробленої інформації на структуровані інсайти для безперешкодної інтеграції та отримання дієвої бізнес-аналітики.",
			en: "We specialize in extracting and processing data from online sources, transforming raw information into structured insights for seamless integration and actionable business intelligence.",
		},
		"telegramTitle": {
			ua: "Розробка ботів",
			en: "Bot Development",
		},
		"telegramDescription": {
			ua: "Ми пропонуємо експертні послуги з розробки ботів для Telegram та інших платформ обміну повідомленнями, виконуючи все, від початкового проектування ботів до розгортання та оптимізації, забезпечуючи безперебійну інтеграцію та підвищену залученість користувачів.",
			en: "We offer expert bot development services for Telegram and other messaging platforms, handling everything from initial bot design to deployment and optimization, ensuring seamless integration and enhanced user engagement.",
		},
		"backendTitle": {
			ua: "Розробка бекенду",
			en: "Backend Development",
		},
		"backendDescription": {
			ua: "Наші послуги з розробки бекенду охоплюють весь спектр від проектування архітектури до впровадження та оптимізації, забезпечуючи надійну функціональність та масштабованість ваших веб-додатків, API та потреб системної інтеграції.",
			en: "Our backend development services encompass the entire spectrum from architecture design to implementation and optimization, ensuring robust functionality and scalability for your web applications, APIs, and systems integration needs.",
		},
		"frontendTitle": {
			ua: "Розробка фронтенду",
			en: "Frontend Development",
		},
		"frontendDescription": {
			ua: "Ми спеціалізуємося на розробці фронтенду, створюючи інтуїтивно зрозумілі інтерфейси та захоплюючий досвід, який резонує з вашою аудиторією, від початкової концепції та дизайну до безшовної інтеграції з внутрішніми системами, гарантуючи, що ваші веб-додатки будуть і візуально привабливими, і функціонально надійними.",
			en: "We specialize in frontend development, crafting intuitive user interfaces and engaging experiences that resonate with your audience, from initial concept and design to seamless integration with backend systems, ensuring your web applications are both visually appealing and functionally robust.",
		},
		"helpDeskTitle": {
			ua: "Help Desk",
			en: "Help Desk",
		},
		"helpDeskShortDescription": {
			ua: "CRM-система для операторів служби підтримки",
			en: "CRM system for support operators",
		},
		"helpDeskFullDescription": {
			ua: "Мета проекту - задовольнити потреби організації в реєстрації, обліку та моніторингу роботи фахівців, які приймають звернення від бенефіціарів. Наша реалізація замінила Google Таблиці, які використовувалися раніше і не задовольняли потреби організації через велике навантаження (реєстрація до 1300 звернень на день у пікові години). Крім того, наша розробка дозволила значно посилити захист даних за рахунок використання системи контролю доступу ABAC.",
			en: "The project's goal is to meet the organization's needs for recording, accounting, and monitoring the work of specialists who receive appeals from beneficiaries. Our implementation has replaced Google Spreadsheets, which were used earlier and did not meet the needs of the organization due to the heavy workload (recording up to 1300 appeals per day at peak times. In addition, our development allowed us to significantly strengthen data protection by using the ABAC access control system.",
		},
		"facilityTitle": {
			ua: "Фасилітація. Громади",
			en: "Facilitation. Communities",
		},
		"facilityShortDescription": {
			ua: "CRM-система для фасилітаторів",
			en: "CRM system for facilitators",
		},
		"facilityFullDescription": {
			ua: "Програмне забезпечення призначене для обліку діяльності та заходів, які проводять фасилітатори. Основним завданням було полегшити облік учасників групових заходів та максимально прискорити його. Крім того, в проекті було використано технології Server-Side Rendering та HTMX, що спростило та прискорило розробку, а також значно зменшило час завантаження сторінок.",
			en: "The software is designed to record the activities and events conducted by facilitators. The main task was to facilitate the accounting of participants in group events and speed it up as much as possible. In addition, the project featured the use of Server-Side Rendering and HTMX technologies, which simplified and accelerated development, as well as significantly reduced page loading time.",
		},
		"autoParserTitle": {
			ua: "AutoParser",
			en: "AutoParser",
		},
		"autoParserShortDescription": {
			ua: "Парсинг та аналіз даних",
			en: "Data parsing and analysis",
		},
		"autoParserFullDescription": {
			ua: "Клієнту потрібен був інструмент для масового моніторингу цін на автозапчастини на декількох маркетплейсах. Для цього ми реалізували веб-портал для приватного використання співробітниками клієнта. Після завантаження excel-файлу зі списком автозапчастин та налаштування процесу парсингу, система збирає інформацію. Завдяки використанню паралелізму та проксі-серверів нам вдалося зробити цей процес дуже швидким. Після збору даних відкидаються позиції, ціни на які, найімовірніше, не відповідають дійсності (за допомогою медіани), дані фільтруються за певними параметрами, після чого зібрана інформація виводиться на окремий екран, звідки її також можна завантажити у вигляді excel-файлу.",
			en: "The customer needed a tool for mass monitoring of prices for auto parts on several marketplaces. For this purpose, we implemented a web portal for private use by the client's employees. After uploading an excel file with a list of auto parts and setting up the parsing process, the system collects information. Thanks to the use of parallelism and proxies, we managed to make this process really fast. After collecting the data, the items whose prices are most likely to be incorrect are discarded (using the median), and the data is filtered by certain parameters, after which the collected information is displayed on a separate screen where it can also be downloaded as an excel file.",
		},
		"messageSent": {
			ua: "Дякуємо! Ваше повідомлення надіслано",
			en: "Thank You! Your message is sent",
		},
		"contactus": {
			ua: "Зв'яжіться з нами",
			en: "Contact Us",
		},
		"name": {
			ua: "Ваше ім'я *",
			en: "Your Name *",
		},
		"email": {
			ua: "Ваш email *",
			en: "Your Email *",
		},
		"phone": {
			ua: "Ваш телефон",
			en: "Your Phone",
		},
		"message": {
			ua: "Ваше повідомлення *",
			en: "Your Message *",
		},
		"sendButton": {
			ua: "Відправити",
			en: "Send Message",
		},
	}

	team = []TeamMember{
		{
			"",
			"Rostyslav Pylypiv",
			"Owner, Golang Developer",
			"",
			"",
			"",
		},
	}
}
