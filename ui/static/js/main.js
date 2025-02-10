let select = document.getElementById("Goods");

if (select) {
	select.addEventListener("change", myFunction);
}



function myFunction() {
	let select = document.getElementById("Goods");
	let textInput = document.getElementById("text");
	let textLabel = document.getElementById("ltext");
	if (select.value !== "") {
		textInput.style = "display: content";
		textLabel.style = "display: content";
	}
}

/*----------------------------------------------------------------------------------------*/

// Получение элемента select HTML с идентификатором "Duos"
const duosSelect = document.getElementById('Duos');

// Добавление обработчика события изменения состояния
duosSelect?.addEventListener('change', (event) => {
	// Проверка значения value
	if (event.target.value !== '') {
		// Значение value не пустое, отправляем GET запрос на сервер
		fetch('/getSecondPerSearch', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({checkFirstEl: event.target.value})
		})
			.then(response => {
				if (response.ok) {
					// Успешный ответ, данные получены с сервера
					return response.json();
				} else {
					// Ошибка при отправке запроса
					console.error("Ошибка при получении данных с сервера.");
				}
			})
			.then(data => {
				// Данные получены и доступны в переменной data
				console.log("Полученные данные: ", data);

				// Получение списка HTML
				const gruzSelect = document.getElementById('Gruz-list');
				for (let i = gruzSelect.options.length - 1; i > 0; i--) {
					gruzSelect.remove(i);
				}

				// Перебор данных и создание элементов списка
				data.forEach(gruz => {
					const option = document.createElement('option');
					option.value = gruz.id;
					option.textContent = `${gruz.name} (ETSNG: ${gruz.etsng}, GNG: ${gruz.gng})`;
					gruzSelect.appendChild(option);
				});
			})
			.catch(error => {
				// Ошибка при выполнении запроса
				console.error("Ошибка при выполнении запроса на сервер: ", error);
			});
	}
});

/*----------------------------------------------------------------------------------------*/
const form = document.getElementById('my-form');
const gruzSelect = document.getElementById('Gruz-list');

form?.addEventListener('submit', (event) => {
	event.preventDefault();
	fetch('/letDuoSearch', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			checkDuos: duosSelect.value,
			checkGruz: gruzSelect.value
		})
	})
		.catch(error => {
			// Ошибка при выполнении запроса
			console.error("Ошибка при выполнении запроса на сервер: ", error);
		})
		.then(response => {
			if (response.ok) {
				// Успешный ответ, данные получены с сервера
				return response.json();
			} else {
				// Ошибка при отправке запроса
				console.error("Ошибка при получении данных с сервера.");
			}
		})
		.then(data => {
			// Данные получены и доступны в переменной data
			console.log("Полученные данные: ", data);

			const tableBody = document.getElementById("app-table-body");
			tableBody.innerHTML = "";

			data.forEach(app => {
				const row = tableBody.insertRow();

				row.insertCell().innerText = app.number_app;
				row.insertCell().innerText = app.regDate_app;
				row.insertCell().innerText = app.status_app;
				row.insertCell().innerText = app.provideDate_app;
				row.insertCell().innerText = app.departureType_app;
				row.insertCell().innerText = app.goods_app;
				row.insertCell().innerText = app.originState_app;
				row.insertCell().innerText = app.enterStation_app;
				row.insertCell().innerText = app.regionDepart_app;
				row.insertCell().innerText = app.roadDepart_app;
				row.insertCell().innerText = app.stationDepart_app;
				row.insertCell().innerText = app.consigner_app;
				row.insertCell().innerText = app.stateDestination_app;
				row.insertCell().innerText = app.exitStation_app;
				row.insertCell().innerText = app.regionDestination_app;
				row.insertCell().innerText = app.roadDestination_app;
				row.insertCell().innerText = app.stationDestination_app;
				row.insertCell().innerText = app.consignee_app;
				row.insertCell().innerText = app.wagonType_app;
				row.insertCell().innerText = app.property_app;
				row.insertCell().innerText = app.wagonOwner_app;
				row.insertCell().innerText = app.payer_app;
				row.insertCell().innerText = app.roadOwner_app;
				row.insertCell().innerText = app.transportManager_app;
				row.insertCell().innerText = app.tonsDeclared_app;
				row.insertCell().innerText = app.tonsAccepted_app;
				row.insertCell().innerText = app.wagonDeclared_app;
				row.insertCell().innerText = app.wagonAccepted_app;
				row.insertCell().innerText = app.filingDate_app;
				row.insertCell().innerText = app.agreementDate_app;
				row.insertCell().innerText = app.approvalDate_app;
				row.insertCell().innerText = app.startDate_app;
				row.insertCell().innerText = app.stopDate_app;
			});
		})
		.catch(error => {
			// Ошибка при выполнении запроса
			console.error("Ошибка при выполнении запроса на сервер: ", error);
		});
});

/*----------------------------------------------------------------------------------------*/

document.querySelectorAll('.menu__switcher > .menu__link').forEach(link => {
	link.addEventListener('click', (event) => {
		event.preventDefault();
		const parent = link.parentElement;
		parent.classList.toggle('active');
	});
});

/*----------------------------------------------------------------------------------------*/
const data = [
	{
		"title": "Главная",
		"url": "/",
		"content": "Главная",
		"snippet": "Главная"
	},
	{
		"title": "Виджеты",
		"url": "/widgets",
		"content": "Виджеты",
		"snippet": "Виджеты"
	},
	{
		"title": "Подвижной состав",
		"url": "/equipment",
		"content": "Подвижной состав",
		"snippet": "Подвижной состав"
	},
	{
		"title": "Галерея",
		"url": "/gallery",
		"content": "Галерея",
		"snippet": "Галерея"
	},
	{
		"title": "Новостная лента",
		"url": "/rss",
		"content": "Новостная лента",
		"snippet": "Новостная лента"
	},
	{
		"title": "Страница с XML данными",
		"url": "/xml",
		"content": "Страница с XML данными",
		"snippet": "Страница с XML данными"
	},
	{
		"title": "О базе данных",
		"url": "/about_db",
		"content": "база данных",
		"snippet": "Данная база данных была создана для быстрого получения и изменения данных о железнодорожных перевозках для поиска новых клиентов и анализа рынка логистических услуг в сфере ж\\д грузоперевозок. Работу выполнил студент группы ИВТ-Б20 - Марков К.А."
	},
	{
		"title": "Ввод данных в бд",
		"url": "/input",
		"content": "Ввод данных в бд",
		"snippet": "Ввод данных в бд"
	},
	{
		"title": "Поиск по двум атрибутам:",
		"url": "/duoSearch",
		"content": "Поиск по двум атрибутам:",
		"snippet": "Поиск по двум атрибутам:"
	},
	{
		"title": "Поиск по одному атрибуту: ",
		"url": "/soloSearch",
		"content": "Поиск по одному атрибуту: ",
		"snippet": "Поиск по одному атрибуту: "
	},
	{
		"title": "Статистика",
		"url": "/stats",
		"content": "Средний уровень грузоперевозок по месяцам:",
		"snippet": "Средний уровень грузоперевозок по месяцам:"
	},
	{
		"title": "Статистика",
		"url": "/stats",
		"content": "Популярные направления для транспортировок: ",
		"snippet": "Популярные направления для транспортировок: "
	},
	{
		"title": "Статистика",
		"url": "/stats",
		"content": "Направления, по которым возят ключевые грузоотправители: ",
		"snippet": "Направления, по которым возят ключевые грузоотправители: "
	}
	,
	{
		"title": "Статистика",
		"url": "/stats",
		"content": "Статистика",
		"snippet": "Статистика"
	}
];

const searchInput2 = document.getElementById('search-input2');
const resultsContainer = document.getElementById('results');

function searchSite() {
  console.log(1212)
	const query = searchInput2.value.toLowerCase();
	resultsContainer.innerHTML = '';


	let results;


	if (query !== '') {
		results = data.filter(page =>
			page.snippet.toLowerCase().includes(query) || page.content.toLowerCase().includes(query)
		);
	}


	results.forEach(result => {
		const resultDiv = document.createElement('div');
		resultDiv.innerHTML = `<h3><a href="${result.url}">${result.title}</a></h3><p>${result.content}</p>`;
		resultsContainer.appendChild(resultDiv);
	});
}

searchInput2.addEventListener('focus', () => {
	resultsContainer.style.display = 'flex';
});

searchInput2.addEventListener('blur', () => {
	setTimeout(() => {
		resultsContainer.style.display = 'none';
	}, 200);
});
searchInput2.addEventListener('input', searchSite);

/*----------------------------------------------------------------------------------------*/

// Получить ссылку на таблицу
const table = document.querySelector('table');
const search_input = document.querySelector('#search-input')

// Событие ввода для поля поиска
if (search_input) {
	search_input.addEventListener('input', function () {
		// Получить значение поиска
		const searchValue = this.value.toLowerCase();

		// Скрыть все строки, которые не содержат значение поиска
		table.querySelectorAll('tbody tr').forEach(function (row) {
			const rowText = row.textContent.toLowerCase();
			if (!rowText.includes(searchValue)) {
				row.style.display = 'none';
			} else {
				row.style.display = '';
			}
		});
	});

// Событие клика для заголовков столбцов для сортировки
	table.querySelectorAll('th').forEach(function (header) {
		header.addEventListener('click', function () {
			// Получить индекс столбца
			const columnIndex = header.cellIndex;

			// Получить все строки в таблице
			const rows = Array.from(table.querySelectorAll('tbody tr'));

			// Отсортировать строки по значению указанного столбца
			rows.sort(function (a, b) {
				const aValue = a.children[columnIndex].textContent;
				const bValue = b.children[columnIndex].textContent;

				return aValue.localeCompare(bValue);
			});

			// Удалить все существующие строки из таблицы
			table.querySelectorAll('tbody tr').forEach(function (row) {
				row.remove();
			});

			// Добавить отсортированные строки обратно в таблицу
			rows.forEach(function (row) {
				table.querySelector('tbody').appendChild(row);
			});
		});
	});
}

/*----------------------------------------------------------------------------------------*/

const buttonContainer = document.getElementById("button-container");
const loginModal = document.getElementById("login-modal");

buttonContainer.addEventListener("click", (event) => {
	event.preventDefault();// Предотвратить выполнение действия по умолчанию (переход по ссылке)
	loginModal.classList.remove("hidden");
});

window.addEventListener("click", (event) => {
	const target = event.target;
	const isInsideSimpleDiv = target.closest('.simple-div') !== null;

	if (target.classList.contains("modal11")) {
		target.classList.add("hidden");
	}
	if (!target.classList.contains("material-symbols-outlined") && !isInsideSimpleDiv && (document.getElementById("profile-list").className === "simple-div")) {
		document.getElementById("profile-list").classList.toggle("hidden");
	}
});
/*----------------------------------------------------------------------------------------*/

const passwordInput = document.getElementById("password");
const confirmPasswordInput = document.getElementById("confirm-password");

// Добавляем обработчик события изменения для поля подтверждения пароля
confirmPasswordInput?.addEventListener("input", () => {
	// Проверяем, совпадают ли пароли
	if (confirmPasswordInput.value !== passwordInput.value) {
		// Пароли не совпадают, показываем сообщение об ошибке
		confirmPasswordInput.setCustomValidity("Пароли не совпадают");
	} else if (confirmPasswordInput.value === passwordInput.value && confirmPasswordInput.validity.customError) {
		// Пароли совпадают, удаляем сообщение об ошибке
		confirmPasswordInput.setCustomValidity("");
	}
});

const reg_form = document.getElementById("reg-form");
const reg_username = document.getElementById("username");
const reg_password = document.getElementById("password");
const reg_email = document.getElementById("email");

reg_username?.addEventListener("change", () => {
	reg_username.setCustomValidity("")
})

reg_email?.addEventListener("change", () => {
	reg_email.setCustomValidity("")
})

reg_form?.addEventListener("submit", async (event) => {
	event.preventDefault();

	const formData = {
		username: reg_username.value,
		password: reg_password.value,
		email: reg_email.value,
	};

	console.log("FormData:", formData);
	fetch("/registration", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(formData)
	})
		.catch(error => {
			console.error("Ошибка при выполнении запроса на сервер: ", error)
		})
		.then(response => {
			if (!response.ok) {
				alert("Something went wrong! Please try again.");
			} else {
				return response.json()
			}
		})
		.then(result => {
			if (result.success === false) {
				const existingUsername = result.existingUsername;
				const existingEmail = result.existingEmail;

				if (existingUsername && reg_username.value === existingUsername) {
					reg_username.setCustomValidity("Пользователь с таким именем уже существует");
				}
				if (existingEmail && reg_email.value === existingEmail) {
					reg_email.setCustomValidity("Пользователь с такой почтой уже существует");
				}
			} else {
				setTimeout(() => {
					window.location.href = "/";
				}, 1000);
				alert("success!");
			}
		})
		.catch(error => {
			// Ошибка при выполнении запроса
			console.error("Ошибка при выполнении запроса на сервер: ", error);
		});
});

/*----------------------------------------------------------------------------------------*/

// Часть передачи запросов авторизации на сервер
const auth_login = document.getElementById("auth-username");
const auth_pass = document.getElementById("auth-password");
const authForm = document.getElementById("authForm")
let token = ''

auth_login.addEventListener("change", () => {
	auth_login.setCustomValidity("")
});

auth_pass.addEventListener("change", () => {
	auth_pass.setCustomValidity("")
});

authForm?.addEventListener("submit", async (event) => {
	event.preventDefault()

	const authData = {
		auth_username: auth_login.value,
		auth_password: auth_pass.value,
	};

	fetch("/authorization", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(authData)
	})
		.catch(error => {
			console.error("Ошибка при выполнении запроса на сервер: ", error)
		})
		.then(authResponse => {
			if (!authResponse.ok) {
				alert("Something went wrong! Please try again.")
			} else {
				token = authResponse.headers.get("Authorization");
				return authResponse.json()
			}
		})
		.then(authResult => {
			const existingUser = authResult.existingUser;
			const existingPass = authResult.existingPass;
			if (existingUser !== "") {
				if (existingPass !== "") {
					setTimeout(() => {
						window.location.reload()
					}, 1000);
					// Сохранить JWT в браузере.
					localStorage.setItem("jwt", token);
				} else {
					auth_pass.setCustomValidity("Неверный пароль")
				}
			} else {
				auth_login.setCustomValidity("Пользователя с таким именем не существует")
			}
		})
		.catch(error => {
			console.error("Ошибка при получении данных с сервера: ", error)
		});
});

// Попытка авторизации (JWT)
// Получить JWT из браузера.
let user_name = ``; // Объявляем user_name в глобальной области видимости

async function getUserInfo() {
	const jwt = localStorage.getItem("jwt");

	if (jwt) {
		try {
			const res = await fetch("/protected", { // Используем await
				method: "POST",
				headers: {
					'Authorization': `${jwt}`
				}
			});

			if (!res.ok) {
				throw new Error(`HTTP error! status: ${res.status}`);
			}

			const data = await res.json();
			document.getElementById("button-container").classList.add('hidden');
			document.getElementById("profile").classList.remove('hidden');
			document.getElementById("user-name-text").innerHTML = `${data.user}`;
			document.getElementById("user-mail-text").innerHTML = `${data.mail}`;
			user_name = `${data.user}`; // Устанавливаем user_name

		} catch (err) {
			console.error('Ошибка при получении данных:', err);
			user_name = ""; // Сбрасываем user_name в случае ошибки
		}
	} else {
		user_name = ""; // Сбрасываем, если нет JWT
	}
	return user_name;
}

const profile = document.getElementById("profile")

profile.addEventListener("click", () => {
	document.getElementById("profile-list").classList.remove('hidden')
});

document.getElementById("logout").addEventListener("click", () => {
	localStorage.removeItem('jwt');
	user_name = ``;
	window.location.reload()
});

//Код для фильтра на новстной странице
document.addEventListener("DOMContentLoaded", function () {
	const newsFilter = document.querySelector('.news-filter');

	if (newsFilter) {
		const links = newsFilter.querySelectorAll('a.clickable');

		if (links.length > 0) {
			links.forEach(link => {
				link.addEventListener('click', function () {
					// Удаляем класс 'active' у всех ссылок
					links.forEach(l => l.classList.remove('active'));
					// Добавляем класс 'active' только к текущему элементу
					this.classList.add('active');

					// Получаем тип новостей из data-атрибута
					const newsType = this.getAttribute('data-type');
					// Вызываем функцию filterNews с нужным типом
					filterNews(newsType);
				});
			});
		}
	}
});

//Передача данных формы добавление новостей на сервер
const addNewBtn = document.getElementById("addNewBtn");
const addNewForm = document.getElementById("addNew-modal");
const addNewTitle = document.getElementById("addNewTitle");
const addNewText = document.getElementById("addNewText");
const addSrcBtn = document.getElementById("addSrcBtn");
const addSrcForm = document.getElementById("addSrc-modal");
const addSrcTitle = document.getElementById("addSrcTitle")

const updateUIBasedOnToken = () => {
	if (user_name !== "") {
		addNewBtn.classList.remove('hidden');
		addSrcBtn.classList.remove('hidden');
	} else {
		addNewBtn.classList.add('hidden');
		addSrcBtn.classList.add('hidden');
	}
};
updateUIBasedOnToken();

console.log('Пользователь: ', user_name);
addNewBtn.addEventListener("click", () => {
	addNewForm.classList.remove("hidden")
});
addSrcBtn.addEventListener("click", () => {
	addSrcForm.classList.remove("hidden")
});

addNewForm?.addEventListener("submit", async (event) => {
	event.preventDefault()

	const addNewData = {
		title: addNewTitle.value,
		text: addNewText.value,
		user_id: user_name,
	};

	fetch("/addNew", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(addNewData)
	})
		.catch(error => {
			console.error("Ошибка при выполнении запроса на сервер: ", error)
		})
		.then(addNewResponse => {
			if (!addNewResponse.ok) {
				alert("Something went wrong! Please try again.");
			} else {
				setTimeout(() => {
					window.location.reload()
				}, 1000);
				alert("Success!");
			}
		})
		.catch(error => {
			console.error("Ошибка при получении данных с сервера: ", error)
		});
});
addSrcForm?.addEventListener("submit", async (event) => {
	event.preventDefault()

	const addNewData = {
		login: user_name,
		url: addSrcTitle.value,
	};

	fetch("/addSource", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(addNewData)
	})
		.catch(error => {
			console.error("Ошибка при выполнении запроса на сервер: ", error)
		})
		.then(addNewResponse => {
			if (!addNewResponse.ok) {
				alert("Something went wrong! Please try again.");
			} else {
				setTimeout(() => {
					window.location.reload()
				}, 1000);
				alert("Success!");
			}
		})
		.catch(error => {
			console.error("Ошибка при получении данных с сервера: ", error)
		});
});

let allNews = [];

async function fetchNews(user_name) {
	try {
		const response = await fetch('/indexNews', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({user: user_name})
		});

		if (!response.ok) {
			const errorData = await response.json(); // попытка получить данные об ошибке
			const errorMessage = errorData.Error ? errorData.Error : response.statusText;
			console.error('Ошибка при запросе:', response.status, errorMessage);
			throw new Error('Network response was not ok: ' + errorMessage);
		}

		let data;
		try {
			data = await response.json(); // Помещаем в try/catch, чтобы поймать ошибки парсинга
		} catch (error) {
			console.error('Ошибка при парсинге JSON:', error);
			throw new Error('Error parsing JSON: ' + error);
		}

		if (!data.feed) { // Проверяем на null
			console.warn('Полученные данные не содержат массива новостей "feed". Возможно сервер вернул ошибку.');
			return []; // Возвращаем пустой массив
		}

		if (!Array.isArray(data.feed)) {
			throw new Error('Полученные данные не содержат массива новостей "feed"');
		}

		allNews = data.feed.map(item => {
			if (item.isLocal) {
				return {
					id: item.local.id,
					title: item.local.title,
					text: item.local.text,
					user: item.local.user,
					date: item.local.resDate,
					type: 'local'
				};
			} else {
				return {
					title: item.global.title,
					link: item.global.link,
					date: item.global.published,
					type: 'global',
					author: item.global.guid || null,
					enclosure: item.global.enclosures || null
				};
			}
		});
		return allNews;
	} catch (error) {
		console.error('Ошибка при загрузке новостей:', error);
		throw error;
	}
}

function displayNews(news) {
	const container = document.getElementById('news-container');
	container.innerHTML = '';

	news.sort((a, b) => new Date(b.date) - new Date(a.date));
	news.forEach(item => {
		const newsItem = document.createElement('div');
		newsItem.className = 'news';

		if (item.type === 'global') {
			newsItem.innerHTML = `
            <a href="${item.link}" style="font-size: 18px">${item.title}</a>
            <div style="flex: 1; display: flex; flex-direction: column; justify-content: space-between;">
               ${item.enclosure?.type ? // Используем опциональную цепочку
				(item.enclosure.type.startsWith("image/") ?
					`<p></p>
						<img style="display: block; margin-left: auto; margin-right: auto; max-width: 100%; height: auto;" src="${item.enclosure.url}"
						alt="Изображение не найдено"/>` :
					``) : ''}
				<div style="display: flex;">
					<a style="text-align: left;" class="no-hover">${item.date}</a>
					${item.author ? `<a style="margin-left: auto;" class="no-hover">${item.author}</a>` : `<a style="margin-left: auto;" class="no-hover">Внешние новости</a>`}
				</div>
			</div>
        `;
		} else {
			newsItem.innerHTML = `
            <a href="/local/${item.id}" style="font-size: 18px">${item.title}</a>
            <div style="flex: 1; display: flex; flex-direction: column; justify-content: space-between;">
				<p></p>
				<div style="display: flex;">
					<a style="text-align: left;" class="no-hover">${item.date}</a>
					<a style="margin-left: auto;" class="no-hover">Новости сайта</a>
				</div>
			</div>
        `;
		}

		container.appendChild(newsItem);
	});
}

function filterNews(type) {
	if (type === 'all') {
		displayNews(allNews);
	} else {
		const filteredNews = allNews.filter(news => news.type === type);
		displayNews(filteredNews);
	}
}

document.addEventListener('DOMContentLoaded', async () => {
	try {
		await getUserInfo(); // Дожидаемся получения информации о пользователе
		const news = await fetchNews(user_name);
		displayNews(news);
		updateUIBasedOnToken();
	} catch (error) {
		console.error('Ошибка при загрузке новостей:', error);
	}
});

const cons_input = document.getElementById('applicationForm');
cons_input?.addEventListener("submit", async (event) => {
	event.preventDefault();
	fetch('/save_application', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			number_app: document.getElementById('Number').value,
			reg_date_app: document.getElementById('Reg_date').value,
			status_app: document.getElementById('Status').value,
			provide_date_app: document.getElementById('Provide_date').value,
			departure_type_app: document.getElementById('Departure_type').value,
			goods_app: document.getElementById('Goods').value,
			origin_state_app: document.getElementById('Origin_state').value,
			enter_station_app: document.getElementById('Enter_station').value,
			region_depart_app: document.getElementById('Region_depart').value,
			road_depart_app: document.getElementById('Road_depart').value,
			station_depart_app: document.getElementById('Station_depart').value,
			consigner_app: document.getElementById('Consigner').value,
			state_destination_app: document.getElementById('State_destination').value,
			exit_station_app: document.getElementById('Exit_station').value,
			region_destination_app: document.getElementById('Region_destination').value,
			road_destination_app: document.getElementById('Road_destination').value,
			station_destination_app: document.getElementById('Station_destination').value,
			consignee_app: document.getElementById('Consignee').value,
			wagon_type_app: document.getElementById('Wagon_type').value,
			property_app: document.getElementById('Property').value,
			wagon_owner_app: document.getElementById('Wagon_owner').value,
			payer_app: document.getElementById('Payer').value,
			road_owner_app: document.getElementById('Road_owner').value,
			transport_manager_app: document.getElementById('Transport_manager').value,
			tons_declared_app: document.getElementById('Tons_declared').value,
			tons_accepted_app: document.getElementById('Tons_accepted').value,
			wagon_declared_app: document.getElementById('Wagon_declared').value,
			wagon_accepted_app: document.getElementById('Wagon_accepted').value,
			filing_date_app: document.getElementById('Filing_date').value,
			agreement_date_app: document.getElementById('Agreement_date').value,
			approval_date_app: document.getElementById('Approval_date').value,
			start_date_app: document.getElementById('Start_date').value,
			stop_date_app: document.getElementById('Stop_date').value,
		})
	})
		.catch(error => {
			// Ошибка при выполнении запроса
			alert("Ошибка при выполнении запроса на сервер: ", error);
		})
		.then(response => {
			if (response.ok) {
				// Успешный ответ, данные получены с сервера
				return response.json();
			} else {
				// Ошибка при отправке запроса
				alert("Ошибка при получении данных с сервера.");
			}
		})
		.then(data => {
			// Данные получены и доступны в переменной data
			console.log("Полученные данные: ", data);
			if (data.success === false){
				setTimeout(() => {
					window.location.href = "/";
				}, 1000);
				alert(`Данные не сохранены. ${data.error}`);
			}else{
				setTimeout(() => {
					window.location.href = "/";
				}, 1000);
				alert(`Заявка принята!`);
			}
		})
});
