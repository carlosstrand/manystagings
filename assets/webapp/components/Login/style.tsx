import styled from 'styled-components';

// Material UI components
import MUIAlert from '@material-ui/lab/Alert';

// Icons
import InputComp from '../../ui/Input';
import LinkComp from '../../ui/Link';

export const Alert = styled(MUIAlert)`
  margin-bottom: 16px;
`;

export const Input = styled(InputComp).attrs({ fullWidth: true })`
  margin-bottom: 16px;
`;

export const ForgotPasswordLink = styled(LinkComp)`
  float: right;
  display: inline-block;
  margin: 16px 0;
`;
