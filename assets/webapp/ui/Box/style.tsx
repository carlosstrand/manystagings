import styled from 'styled-components';

export const Wrapper = styled.div<{ padding?: number }>`
  position: relative;
  border-radius: 8px;
  background-color: #ffffff;
  padding: ${props => props.padding === undefined ? 32 : props.padding}px;
`;

export const SpinnerOverlay = styled.div<{ show: boolean }>`
  transition: opacity 300ms ease-in;
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  background-color: rgba(255, 255, 255, 0.7);
  border-radius: 8px;
  height: ${(props) => (props.show ? '100%' : '0')};
  opacity: ${(props) => (props.show ? '1' : '0')};
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
`;
