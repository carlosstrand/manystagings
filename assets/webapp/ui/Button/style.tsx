import styled from 'styled-components';
import MUIButton, { ButtonProps } from '@material-ui/core/Button';
import MUICircularProgress from '@material-ui/core/CircularProgress';

export const Spinner = styled(MUICircularProgress).attrs({ size: 26, thickness: 5 })`
  color: #fff;
  padding: 4px;
`;

export const StyledButton = styled(MUIButton)<ButtonProps & { loading?: boolean }>`
  box-shadow: none;
  border-radius: 8px;
  &:hover {
    box-shadow: none;
  }
  ${(props) =>
    props.loading &&
    `
    cursor: auto;
    pointer-events: none;
    &:hover, &:active {
      box-shadow: none;
    }
    .MuiTouchRipple-root {
      display: none;
    }
  `}
  ${(props) =>
    (!props.color || props.color === 'default') &&
    `
    border: 1px solid #666;
    background-color: white;
    color: #666;
    &:hover {
      box-shadow: none;
      background-color: #e9e9e9;
      color: #666;
    }
    .MuiTouchRipple-root {
      color: #ccc;
    }
  `}
`;
