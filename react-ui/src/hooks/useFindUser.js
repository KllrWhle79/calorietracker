import { useState, useEffect } from 'react';
import axios from 'axios'; 

export default function useFindUser() {
    const [user, setUser] = useState(null);
    const [isLoading, setLoading] = useState(true);

    useEffect(() =>{
        setLoading(false);
    }, []);
    
    return {
        user,
        setUser,
        isLoading
    }
}