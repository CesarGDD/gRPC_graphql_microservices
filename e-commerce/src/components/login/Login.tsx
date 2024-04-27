import { Link, useNavigate } from 'react-router-dom';
import React, { useEffect } from 'react';
import './Login.css'
import { useSelector } from 'react-redux';
import { selectUser } from '../../store/users/userSlice';

const Login = () => {
    const user = useSelector(selectUser);
    let navigate = useNavigate();
    
    useEffect(() => {
        if (user) {
          navigate('/');
        }
      }, [user, navigate]);

    return (
        <div className="login" >
            {/* {user ? navigate('/') : null} */}
            <Link className="login__link" to='/' >
            <h1 className="header__logo" > Lux Watch </h1>
            </Link>
            <div className="login__container">
                {/* @ts-ignore */}
                <button className="login__signInButton" onClick={signIn} >Sign-In</button>
                <p>By signing-in you agree to the Lux Watch Conditions of use & Sale. Please see our privacy Notice, our cookies notice and our Interest-Based Ads Notice</p>
            </div>
        </div>
    )
}

export default Login;
