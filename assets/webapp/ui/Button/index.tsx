import { ButtonProps as MUIButtonProps } from '@material-ui/core/Button';
import React from 'react';
import { StyledButton, Spinner } from './style';

interface CustomButtonProps {
  loading?: boolean;
}

export type ButtonProps = CustomButtonProps & MUIButtonProps;

const Button = (props: ButtonProps) => {
  const bprops = { ...props };
  if (bprops.loading) {
    bprops.children = <Spinner />;
    bprops.onClick = () => {};
  }
  return <StyledButton {...bprops} />;
};

Button.defaultProps = {
  variant: 'contained',
  loading: false,
};

export default Button;
