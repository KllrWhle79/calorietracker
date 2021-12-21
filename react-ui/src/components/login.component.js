import React, {useState} from "react";
import PropTypes from "prop-types";

async function loginUser(credentials) {
    return fetch('http://localhost:8000/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
    })
        .then(data => data.json())
}

export default function Login({setToken}) {
    const [username, setUserName] = useState();
    const [password, setPassword] = useState();

    const handleSubmit = async e => {
        e.preventDefault();
        const token = await loginUser({
            "user_name": username,
            "password": password
        });
        setToken(token);
    }

    return (
        <form>
            <h3>Sign In</h3>

            <div className="form-group">
                <label>Username</label>
                <input type="text" className="form-control" placeholder="Enter username" onChange={e => setUserName(e.target.value)}/>
            </div>

            <div className="form-group">
                <label>Password</label>
                <input type="password" className="form-control" placeholder="Enter password" onChange={e => setPassword(e.target.value)}/>
            </div>

            {/*<div className="form-group">*/}
            {/*    <div className="custom-control custom-checkbox">*/}
            {/*        <input type="checkbox" className="custom-control-input" id="customCheck1" />*/}
            {/*        <label className="custom-control-label" htmlFor="customCheck1">Remember me</label>*/}
            {/*    </div>*/}
            {/*</div>*/}

            <button type="submit" className="btn btn-primary btn-block" onClick={handleSubmit}>Submit</button>
            <p className="forgot-password text-right">
                Forgot <a href="#">password?</a>
            </p>
        </form>
    );
}

Login.propTypes = {
    setToken: PropTypes.func.isRequired
}