import styled from 'styled-components';
import Box from '../Box';
import { StyledButton } from '../Button/style';
import { StyledTextField } from '../Input/style';


export const Wrapper = styled.div`
  
`;

export const Dialogue = styled(Box)`
  min-width: 320px;
  max-width: 480px;
`;

export const Header = styled.div`
  padding: 8px 32px;
  border-bottom: 1px solid #e9e9e9;
`;

export const Body = styled.div`
  padding: 32px;
  ${StyledTextField} {
    margin-bottom: 20px;
  }
`;

export const Footer = styled.div`
  padding: 16px 32px;
  border-top: 1px solid #e9e9e9;
  clear: both;
  ${StyledButton} {
    margin-left: 8px;
  }
`;
