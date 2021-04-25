import styled from 'styled-components';
import BoxComp from '../../ui/Box';

export const Wrapper = styled.div`
  width: 100vw;
  height: 100vh;
  background-color: #0C2461;
  padding-top: 64px;
`;

export const LogoWrapper = styled.div`
  with: 100%;
  text-align: center;
  margin: 64px;
`;

export const LogoImg = styled.img.attrs({ src: '/public/img/logo-dark.svg' })`
  height: 32px;
`;

export const Box = styled(BoxComp)`
  margin: 0 auto;
  max-width: 480px;
  width: 100%;
`;
