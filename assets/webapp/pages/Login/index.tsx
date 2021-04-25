import React from 'react';
import { useState } from 'react';
import CenteredBoxLayoyt from '../../layouts/CenteredBox';
import Login from '../../components/Login';

const LoginPage = () => {
  const [loading, setLoading] = useState(false);
  return (
    <CenteredBoxLayoyt withLogo loading={loading}>
      <Login setLoading={setLoading} />
    </CenteredBoxLayoyt>
  );
};

export default LoginPage;
