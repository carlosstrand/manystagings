import React from 'react';
import CircularProgress from '@material-ui/core/CircularProgress';
import { Wrapper, SpinnerOverlay } from './style';

interface BoxProps {
  className?: string;
  children?: React.ReactNode;
  loading?: boolean;
  padding?: number;
}

const Box = (props: BoxProps) => {
  const { className, children, loading, padding } = props;
  return (
    <Wrapper className={className} padding={padding}>
      {children}
      <SpinnerOverlay show={loading || false}>
        <CircularProgress size={32} />
      </SpinnerOverlay>
    </Wrapper>
  );
};

Box.defaultProps = {
  className: '',
  children: null,
  loading: false,
};

export default Box;
