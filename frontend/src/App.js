
import './App.css';
import React, { useState, useEffect } from 'react';
import { BrowserRouter, Routes, Route, useParams, useNavigate } from 'react-router-dom';

function Table() {
  const navigate = useNavigate()
  const [users, setUsers] = useState([]);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await fetch('api/api/children', {
          method: "GET"
        });
        const data = await response.json();
        console.log(data)
        setUsers(data.children);
        console.log(users[0])
      } catch (error) {
        console.error('Ошибка при получении пользователей:', error);
      }
    };

    fetchUsers();
  }, []);
  console.log(users)
  return (
    <div>
      <h1>Таблица детей</h1>
      <ul className='table'>
        <li className='title'>
            <h3>Имя ребенка</h3>
            <h3>Фамилия ребенка</h3>
            <h3>Возраст ребенка</h3>
            <h3>Желаемый подарок</h3>
            <h3>Статус</h3>
            <h3>Ссылка на форму</h3>
        </li>
        {users ? users.map((user) => (
          <li key={user.id} className='child'>
            <h3>{user.FirstName}</h3>
            <h3>{user.LastName}</h3>
            <h3>{user.Age > 0 ? user.Age : ""}</h3>
            <a onClick={() => {navigate(`/letter/${user.PhotoUrl.replace(/\//g, "*slash")}`); console.log(user.PhotoUrl.replace(/\//g, "*slash"))}}>{user.Gift}</a>
            <h3>{user.IsTaken ? "занят" : "не занят"}</h3>
            <h3 onClick={() => {navigate(!(user.IsTaken) ? `/form/${user.Id}` : "")}}>Подарить подарок</h3>
          </li>
        )) : ""}
      </ul>
      
      <button className='button' onClick={() => navigate("/admin")}>Страница для админов</button>
    </div>
  );
}



function Form() {
  const navigate = useNavigate()
  const [ok, setOk] = useState(false)
  const [name, setName] = useState("")
  const [gift, setGift] = useState("")
  const id = useParams()
  console.log(id.id)
  return (
    <>
    <h1>Для того, чтобы выбрать ребенка для подарка, заполните форму</h1>
    <input id='email' placeholder='Введите почту'/>
    <input id='firstname' placeholder='Введите имя'/>
    <input id='lastname' placeholder='Введите фамилию'/>
    <input id='tg' placeholder='Введите username в Telegram'/>
    <input id='phone' placeholder='Введите номер телефона'/>
    <input id='class' placeholder='Введите класс'/>
    <button onClick={SendCode}>Подтвердить почту</button>
    <h2>На почту был отправлен код. Также проверьте папку "Спам"</h2>
    <input id='code' placeholder='Введите код, отправленный на вашу почту'/>
    <button onClick={() => TakeChild(id)}>Выбрать ребенка</button>
    <h1>{ok ? `Вы дарите ${gift} ребенку ${name}` : ""}</h1>
    <button onClick={() => navigate("/children")}>Вернуться на главную страницу</button>
    </>
  )
  async function TakeChild(id) {
    const email = document.getElementById('email').value
    const code = document.getElementById("code").value
    const firstName = document.getElementById("firstname").value
    const lastName = document.getElementById("lastname").value
    const tg = document.getElementById("tg").value
    const phone = document.getElementById("phone").value
    const userClass = document.getElementById("class").value
    console.log(typeof code)
    if (email == "" || code == "" || firstName == "" || lastName == "" || tg == "" || phone == "" || userClass == "") {
      alert("Заполните все поля")
      return
    }
    const response = await fetch(`api/api/children/${id.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({email: email, child_id: id.id, code: code, first_name: firstName, last_name: lastName, telegram: tg, phone_number: phone, class: userClass})
    });
    const data = await response.json();
    if (!(data?.child_fullname != "")) {
      if (data?.error != "") {
        alert("ребенок оказался занятым, либо данные введены некорректно")
      }
    } else {
      setOk(true)
      setGift(data.child_gift)
      setName(data.child_fullname)
    }
  }
  async function SendCode() {
    const email = document.getElementById('email').value
    const response = await fetch('api/api/codes', {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({email: email})
    });
    const data = await response.json();
    if (!(data?.status == "ok")) {
      alert("некорректная почта, либо ребенок уже занят")
    }
  }
}

function About() {
  const [page, setPage] = useState(1)
  const navigate = useNavigate()
  if (page == 1) {
    return(
      <>
      <div style={{alignContent: 'center', alignItems: 'center', width: '80vw', marginLeft: '10vw', marginTop: '10vh'}}>
        <h1 style={{textAlign: 'center'}}>Здравствуй!</h1>
        <p style={{textAlign: 'center'}}>
        Меня зовут Гаркуша Юля, и я учусь в 9 классе школы Летово.
    
    Родом я из Новосибирска, и у меня много увлечений; например, в свободное время я люблю
    плавать, лепить из глины, а также готовить. Однако помимо хобби, у меня с самого детства есть
    непреодолимое желание помогать людям. Так, любимой традицией нашей семьи долгие годы
    
    является участие в новогодней акции от известного в моем регионе
    благотворительного фонда «Солнечный город».
    
    Данное мероприятие всегда вызывает у меня лишь положительные эмоции, объединяет нашу
    
    семью в волшебную предновогоднюю пору.
        </p>
        <h2 style={{textAlign: 'center'}}>Но о чем же я собираюсь рассказать сейчас?</h2>
        <p style={{textAlign: 'center'}}>
        Я уверена, что многие наслышаны про персональные проекты, являющиеся обязательной
    частью учебного года каждого девятиклассника Летово. Узнав про мои интересы, вы сможете
    
    легко догадаться, что мой проект связан с благотворительностью.
    
    Главной его целью является организовать новогоднее мероприятие для детей из детского
    дома в области близ Новосибирска, где каждый из них получит подарок, упомянутый в своем
    
    новогоднем письме, и почувствует атмосферу волшебного праздника.
        </p>
        <h2 style={{color: "red", textAlign: 'center'}}>Но мой проект невозможно будет осуществить без вашей помощи!</h2>
        <h1 style={{textAlign: 'center'}}>Итак, в чем суть моей акции?</h1>
        <p style={{textAlign: 'center'}}>
        Идея в том, что каждый из вас сможет сыграть роль рождественского эльфа, <p style={{color: "red", fontWeight: 'bold'}}>подготовив
        подарок для ребенка из детского дома, о котором тот писал в своем новогоднем письме</p>, чтобы
    в день проведения праздничного мероприятия я (конечно же, от лица волшебных эльфов!)
    
    подарила подготовленные вами подарки детям.
    
    Кроме того, обязательным условием является <p style={{color: "red", fontWeight: 'bold', textAlign: "center"}}>написать ребенку мотивационное письмо</p>, в
    
    котором вы
    
    поддержите его интересы (данная информация будет указана в его письме, о чем я расскажу
    совсем скоро), а может захотите и дальше общаться с ним, поддерживая дружескую переписку.
    Написать письмо - очень доброе дело, которое точно принесет положительные эмоции ребенку,
    однако, чтобы оно вызвало у вас больше радости и получилось ещё более искренним, <p style={{color: "red", fontWeight: 'bold'}}>за его
    написание вам зачислятся определенные балы в диплом Летово!</p>
        </p>
        
      </div>
      <button className='button' onClick={() => setPage(page + 1)}>Идти дальше</button>
      </>
    )
  }
  if (page == 2) {
    return (
      <>
      <div style={{alignContent: 'center', alignItems: 'center', width: '80vw', marginLeft: '10vw', marginTop: '10vh'}}>
        <h2>Таким образом, задачи у участников акции будут несложные и очень интересные:</h2>
        <p>
          1. Ознакомившись с письмами детей, выбрать ребенка, которому участник акции (или небольшая группа учеников) захочет подарить подарок и заполнить небольшую форму, тем самым «забронировав» выбранный подарок (желательно сделать это до …)
        </p>
        <p>
          2. Подготовить подарок и упаковать его в тот вид, в котором он будет подарен ребенку.
        </p>
        <p>
          3. Написать письмо, сфотографировать его так, чтобы было видно, что работа принадлежит вам. *если ученики решили объединиться в группу, чтобы подарить один подарок от всех её участников, то письмо может быть одним на всех
        </p>
        <p>
          4. Надежно упаковав подарок и письмо (важно, чтобы письмо было в конверте, но внутрь подарка его складывать не нужно), отправить его одной из транспортных компаний по указанным данным получателя (данные будут представлены непосредственно заинтересовавшимся в акции) до ....
        </p>
        <p>
          5. Внести трек номер и транспортную компанию в ту же форму (из первого пункта).
        </p>
        <p>
          6. Гордиться собой и в скором времени получить фотографии детей с дня проведения праздника, новогоднее письмо выбранного вами ребенка на память, а также балы в диплом Летово!
        </p>
        <p>
          По всем возникшим вопросам вы можете писать мне в ТГ (@Julis0009)
        </p>
        <p>
          Ознакомиться с информацией о детях и стать участником акции вы сможете прямо сейчас, перейдя к таблице и заполнив форму:
        </p>
      </div>
      <button className='button' onClick={() => navigate("/children")}>Стать рождественским эльфом</button>
      </>
    )
  }
  
}

function Admin() {
  const [code, setCode] = useState(false)
  const [users, setUsers] = useState([]);
  return(
      <>
      <h2>Напишите свою почту</h2>
      <input id='email' placeholder='Почта' />
      <button onClick={SendCode}>Отправить код</button>
      <h2>Ведите код, отправленный на вашу почту. Также проверьте папку спам</h2>
      <input id='code' placeholder='Код'/>
      <button onClick={fetchUsers}>Подтвердить почту</button>
      <div>
      <h1>Таблица пользователей, занявших детей</h1>
      <ul className='table'>
        {users ? (
          <li className='title'>
            <h3 className='h3'>Почта</h3>
            {/* <h3 className='h3'>Имя</h3> */}
            <h3 className='h3'>Фамилия</h3>
            <h3 className='h3'>Телефон</h3>
            <h3 className='h3'>Телеграм</h3>
            <h3 className='h3'>Класс</h3>
            <h3 className='h3'>Ребенок</h3>
          </li>
        ) : ""}
        
        {users ? users.map((user) => (
          <li key={user.id} className='child'>
            <h3 className='h3'>{user.UserEmail}</h3>
            {/* <h3 className='h3'>{user.UserFirstName}</h3> */}
            <h3 className='h3'>{user.UserLastName}</h3>
            <h3 className='h3'>{user.UserPhoneNumber}</h3>
            <h3 className='h3'>{user.UserTelegram}</h3>
            <h3 className='h3'>{user.UserClass}</h3>
            <h3 className='h3'>{user.FirstName + " " + user.LastName}</h3>
          </li>
        )) : ""}
      </ul>
    </div>
      </>
  )

  async function SendCode() {
    const email = document.getElementById('email').value
    const response = await fetch('api/api/codes', {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({email: email})
    });
    const data = await response.json();
    if (!(data?.status == "ok")) {
      alert("некорректная почта, либо ребенок уже занят")
    } else {
      setCode(true)
    }
  }

  async function fetchUsers() {
    try {
      const response = await fetch('api/api/admin/children', {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({email: document.getElementById("email").value, code: document.getElementById("code").value})
      });
      const data = await response.json();
      console.log(data)
      setUsers(data.children);
      console.log(users[0])
    } catch (error) {
      console.error('Ошибка при получении пользователей:', error);
      alert("неверный код, либо вас нет в списке админов")
    }
  }
}

function Letter() {
  const url = useParams()
  const navigate = useNavigate()
  console.log(url.url)
  return (
    <>
    <button onClick={() => navigate("/children")}>Назад</button>
    <div style={{display: 'flex', alignContent: 'center', alignItems: 'center', justifyContent: 'center'}}>
      <img style={{height: '50vh', borderRadius: '10px'}} src={url.url.replace(/\*slash/g, "/")}/>
    </div>
    
    </>
    
  )
}

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path='/' element={<About/>}/>
        <Route exact path='/children' element={<Table/>}/>
        <Route exact path='/form/:id' element={<Form/>}/>
        <Route exact path='/admin' element={<Admin/>}/>
        <Route exact path='/letter/:url' element={<Letter/>}/>
      </Routes>
    </BrowserRouter>
  )
  
}

export default App;

