import React from 'react';
import { Wrapper, LogoWrapper, LogoImg, Box } from './style';

interface CenteredBoxLayoytProps {
  children: React.ReactNode;
  withLogo?: boolean;
  loading?: boolean;
}

const CenteredBoxLayoyt = (props: CenteredBoxLayoytProps) => {
  const { children, withLogo, loading } = props;
  return (
    <Wrapper>
      {withLogo && (
        <LogoWrapper>
          <LogoImg />
        </LogoWrapper>
      )}
      <Box loading={loading}>{children}</Box>
    </Wrapper>
  );
};

CenteredBoxLayoyt.defaultProps = {
  loading: false,
  withLogo: false,
};

export default CenteredBoxLayoyt;
