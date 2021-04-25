import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import Person from '@material-ui/icons/Person';

import PasswordInput from '../../ui/PasswordInput';
import Button from '../../ui/Button';

import useAuth from '../../hooks/useAuth';

import { Input, ForgotPasswordLink, Alert } from './style';

type LoginProps = {
  className?: string;
  setLoading?: (value: boolean) => void;
};

const Login = ({ className, setLoading }: LoginProps) => {
  const history = useHistory();
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const { mutate, isLoading, isError, error } = useAuth({
    onSuccess: (res) => {
      localStorage.setItem('token', JSON.stringify(res.token));
      history.push('/select-environment');
    },
  });
  useEffect(() => {
    if (setLoading) {
      setLoading(isLoading);
    }
  }, [isLoading]);

  const onSubmit = (e: React.FormEvent<EventTarget>) => {
    e.preventDefault();
    mutate({ username, password });
  };

  return (
    <div className={className}>
      {isError && (
        <Alert severity="error">
          {error?.response.status === 401
            ? 'Username or password is invalid. Please try again'
            : 'Internal Server Error. please try again later'}
        </Alert>
      )}
      <form onSubmit={onSubmit}>
        <Input
          placeholder="Username"
          type="text"
          leftIcon={<Person />}
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <PasswordInput
          placeholder="Password"
          fullWidth
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <ForgotPasswordLink to="/forgot-password">Forgot your password?</ForgotPasswordLink>
        <Button type="submit" color="primary" fullWidth size="large">
          Login
        </Button>
      </form>
    </div>
  );
};

Login.defaultProps = {
  className: '',
  setLoading: () => null,
};

export default Login;
