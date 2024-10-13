import { Link, useNavigate } from 'react-router-dom';
import React, { useEffect, useState } from 'react';
import './Login.css';
import { useSelector, useDispatch } from 'react-redux';
import { selectUser, login as loginAction } from '../../store/users/userSlice';
import { useLoginMutation } from '../../graphql/generated/graphql';

const Login = () => {
    const user = useSelector(selectUser);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [loginError, setLoginError] = useState<string | null>(null);
    const navigate = useNavigate();
    const dispatch = useDispatch();
    const [{ fetching, error }, login] = useLoginMutation();

    useEffect(() => {
        if (user) {
            navigate('/');
        }
    }, [user, navigate]);

    // Make this function async to use `await` inside it
    const handleLogin = async () => {
        try {
            const response = await login({ input: { username, password } });

            if (response.data?.login.token) {
                // Dispatch the login action to Redux store
                dispatch(loginAction({
                    userId: response.data.login.user.userId,
                    username: response.data.login.user.username,
                    token: response.data.login.token,
                }));

                // Redirect to home page after successful login
                navigate('/');
            } else {
                // Handle login failure
                setLoginError('Invalid login credentials.');
            }
        } catch (err) {
            console.error('Login failed:', err);
            setLoginError('An error occurred. Please try again.');
        }
    };

    return (
        <div className="login">
            <Link className="login__link" to="/">
                <h1 className="header__logo">E-commerce</h1>
            </Link>
            <div className="login__container">
                <h2>Sign In</h2>
                <input
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    placeholder="Username"
                    className="login__input"
                />
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="Password"
                    className="login__input"
                />
                {loginError && <p className="login__error">{loginError}</p>}
                <button
                    className="login__signInButton"
                    onClick={handleLogin}
                    disabled={fetching}
                >
                    {fetching ? 'Signing in...' : 'Sign-In'}
                </button>
                {error && <p className="login__error">Error logging in. Please try again.</p>}
                <p>
                    By signing-in you agree to the E-commerce Conditions of Use & Sale. Please see our Privacy Notice, Cookies Notice, and Interest-Based Ads Notice.
                </p>
            </div>
        </div>
    );
};

export default Login;
