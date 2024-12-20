import profilePicture from './assets/hello.png'

function Card() {
    return (
        <div className="card">
            <img src={profilePicture} alt="profile picture"></img>
            <h2>Poller</h2>
            <p>Aparna is creating this poller app</p>

        </div>
    );
}

export default Card;