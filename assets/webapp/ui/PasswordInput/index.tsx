import { useState } from 'react';
import IconButton from '@material-ui/core/IconButton';
import Lock from '@material-ui/icons/Lock';
import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';

import Input, { InputProps } from '../Input';
import React from 'react';

const PasswordInput = (props: InputProps) => {
  const [showPassword, setShowPassword] = useState(false);
  return (
    <Input
      type={showPassword ? 'text' : 'password'}
      leftIcon={<Lock />}
      rightAction={
        <IconButton onClick={() => setShowPassword(!showPassword)}>
          {showPassword ? <Visibility /> : <VisibilityOff />}
        </IconButton>
      }
      {...props}
    />
  );
};

export default PasswordInput;
