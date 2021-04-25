import styled from 'styled-components';
import TextField from '@material-ui/core/TextField';
import InputAdornment from '@material-ui/core/InputAdornment';

export const StyledTextField = styled(TextField)`
  fieldset {
    border-radius: 8px;
  }
`;

export const LeftIconWrapper = styled(InputAdornment)<{ isActive: boolean }>`
  svg {
    color: ${(props) => (props.isActive ? '#0C2461' : '#a0a0a0')};
  }
`;
