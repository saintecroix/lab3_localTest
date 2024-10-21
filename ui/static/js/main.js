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
			body: JSON.stringify({ checkFirstEl: event.target.value })
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

authForm?.addEventListener("submit", async (event) =>{
	event.preventDefault()

	const authData = {
		auth_username: auth_login.value,
		auth_password: auth_pass.value,
	};
	console.log(authData);

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
			if (!authResponse.ok){
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
						window.location.href = "/";
					}, 1000);
					// Сохранить JWT в браузере.
					localStorage.setItem("jwt", token);
				}else {
					auth_pass.setCustomValidity("Неверный пароль")
				}
			}else {
				auth_login.setCustomValidity("Пользователя с таким именем не существует")
			}
		})
		.catch(error => {
			console.error("Ошибка при получении данных с сервера: ", error)
		});
});

// Попытка авторизации (JWT)
// Получить JWT из браузера.
const jwt = localStorage.getItem("jwt");
if (jwt) {
	// Отправить запрос на защищенный маршрут с включенным JWT в заголовке авторизации.
	fetch("/protected", {
		headers: {
			'Authorization': `${jwt}`
		}
	})
		.then(res => res.json())
		.then(data => {
			document.getElementById("button-container").classList.add('hidden')
			document.getElementById("profile").classList.remove('hidden')
			document.getElementById("user-name-text").innerHTML = ` ${data.user}`
			document.getElementById("user-mail-text").innerHTML = `${data.mail}`
		})
		.catch(err => console.error(err));
}

const profile = document.getElementById("profile")

profile.addEventListener("click", () => {
	document.getElementById("profile-list").classList.remove('hidden')
});

document.getElementById("logout").addEventListener("click", () => {
	localStorage.removeItem('jwt');
	window.location.href = "/";
});

//Код для фильтра на новстной странице
document.addEventListener("DOMContentLoaded", function () {
	// Проверяем, есть ли элемент с классом 'news-filter' на странице
	const newsFilter = document.querySelector('.news-filter');

	if (newsFilter) {
		// Находим все элементы <a> внутри .news-filter
		const links = newsFilter.querySelectorAll('a.clickable');

		// Проверяем, есть ли элементы <a> с классом 'clickable'
		if (links.length > 0) {
			links.forEach(link => {
				link.addEventListener('click', function () {
					// Удаляем класс 'active' у всех ссылок
					links.forEach(l => l.classList.remove('active'));
					// Добавляем класс 'active' только к текущему элементу
					this.classList.add('active');
				});
			});
		}
	}
});