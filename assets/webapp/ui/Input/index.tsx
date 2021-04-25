import React, { useState } from 'react';
import { TextFieldProps } from '@material-ui/core/TextField';
import { OutlinedInputProps } from '@material-ui/core/OutlinedInput';
import InputAdornment from '@material-ui/core/InputAdornment';

import { StyledTextField, LeftIconWrapper } from './style';

interface CustomProps {
  leftIcon?: React.ReactNode;
  rightAction?: React.ReactNode;
}

export type InputProps = CustomProps & TextFieldProps;

const Input = (props: InputProps) => {
  const [active, setActive] = useState(false);
  const { leftIcon, rightAction } = props;
  const inputProps: OutlinedInputProps = {};
  if (leftIcon) {
    inputProps.startAdornment = (
      <LeftIconWrapper position="start" isActive={active}>
        {leftIcon}
      </LeftIconWrapper>
    );
  }
  if (rightAction) {
    inputProps.endAdornment = <InputAdornment position="end">{rightAction}</InputAdornment>;
  }
  return (
    <StyledTextField
      InputProps={inputProps}
      {...props}
      onFocus={(e) => {
        setActive(true);
        if (props.onFocus) {
          props.onFocus(e);
        }
      }}
      onBlur={(e) => {
        setActive(false);
        if (props.onBlur) {
          props.onBlur(e);
        }
      }}
    />
  );
};

Input.defaultProps = {
  leftIcon: null,
  variant: 'outlined',
};

export default Input;
