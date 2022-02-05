import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import {useContext} from "react";
import {UserContext} from "./UserContext";

export default function useLogout() {
    let navigate = useNavigate();
    const {setUser} = useContext(UserContext);

    const logoutUser = () => {
        setUser(null);
        navigate('/');
    }

    return {
        logoutUser
    }

}