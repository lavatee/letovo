import React from 'react';

function By() {
    return(
        <>
      <div style={{display: 'flex', alignContent: 'center', alignItems: 'center', justifyContent: 'center', backgroundColor: 'white'}}>
        <div>
            <h1>Gryaznov Alexander</h1>
            <h2>Backend Engineer</h2>
            <h2>Telegram: @gryaznybilly</h2>
            <h2>Github: https://github.com/lavatee</h2>
        </div>
        
        <img src='https://raw.githubusercontent.com/lavatee/facepalm/refs/heads/main/img/photo_5237947810337384444_y.jpg' style={{height: '50vh', marginLeft: '10vw'}}/>
      </div>
      <div style={{display: 'flex', alignContent: 'center', alignItems: 'center', justifyContent: 'space-around', backgroundColor: 'white'}}>
      <h4>Skills:</h4>
      <div style={{marginLeft: '1vw'}}>
          <p>Golang</p>
          <p>SQL</p>
          <p>WebSockets</p>
          <p>PostgreSQL</p>
          <p>Microservices</p>
          <p>gRPC</p>
          <p>Kafka</p>
          <p>MongoDB</p>
          <p>Redis</p>
          <p>HTML/CSS/JS/REACT</p>
          <p>Linux</p>
          <p>CI/CD</p>
      </div>
      <h4 style={{marginLeft: '5vw'}}>Experience:</h4>
      <div style={{marginLeft: '1vw'}}>
          
          <div>
                <b>Pet-Projects:</b>
                <p></p>
              <a href='https://github.com/lavatee/mafia'>Online Mafia</a>
              <p></p>
              <a href='https://github.com/lavatee/messenger'>Messenger</a>
          </div>
          <p style={{width: '40vw'}}>Now I'm getting experience working with microservices using gRPC and Kafka to communicate them. I also have few projects in Gitlab that I'm currently developing</p>
      </div>
    </div>
    </>
    )
  }

export default By