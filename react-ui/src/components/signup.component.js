import React, {useState} from "react";

async function createNewUser(newUser) {
    return fetch('http://localhost:8000/user', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(newUser)
    })
        .then(data => data.json())
}

export default function SignUp() {
    const [username, setUsername] = useState();
    const [emailAddr, setEmailAddr] = useState();
    const [password, setPassword] = useState();

    const handleSubmit = async e => {
        e.preventDefault();
        const user = await createNewUser({
            "user_name": username,
            "email_addr": emailAddr,
            "password": password,
            "admin": false
        })
            .then(data => data.json())
    }

    return (
        <form>
            <h3>Sign Up</h3>

            <div className="form-group">
                <label>Username</label>
                <input type="text" className="form-control" placeholder="Username"
                       onChange={e => setUsername(e.target.value)}/>
            </div>

            <div className="form-group">
                <label>Email address</label>
                <input type="email" className="form-control" placeholder="Enter email"
                       onChange={e => setEmailAddr(e.target.value)}/>
            </div>

            <div className="form-group">
                <label>Password</label>
                <input type="password" className="form-control" placeholder="Enter password"
                       onChange={e => setPassword(e.target.value)}/>
            </div>

            <button type="submit" className="btn btn-primary btn-block" onClick={handleSubmit}>Sign Up</button>
            <p className="forgot-password text-right">
                Already registered <a href="/sign-in">sign in?</a>
            </p>
        </form>
    );
}