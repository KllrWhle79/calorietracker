import {useState, useContext} from 'react';
import {useNavigate} from 'react-router-dom';
import axios from 'axios';
import {UserContext} from './UserContext';

export default function useAuth() {
    let navigate = useNavigate();
    const {setUser} = useContext(UserContext);
    const [error, setError] = useState(null);

    //register user  
    const registerUser = async (data) => {
        const {username, email, password, firstname, calorie_max} = data;
        return axios.put(`http://localhost:8000/user`, {
            "user_name": username,
            "email_addr": email,
            "first_name": firstname,
            password,
            "admin": false,
            "calorie_max": Number(calorie_max)
        }).then(async () => {
            await loginUser({"username": username, "password": password});
        })
            .catch((err) => {
                return setError(err.response.data);
            })
    };

    //login user 
    const loginUser = async (data) => {
        const {username, password} = data;
        return axios.post('http://localhost:8000/login', {
            "user_name": username,
            password,
        }).then(response => {
            setUser(response.data);
            navigate('/home');
        }).catch((err) => {
            setError(err.response.data);
        })
    };

    return {
        registerUser,
        loginUser,
        error
    }
}
