import { forwardRef } from 'react';
import { Link as RouterLink, LinkProps as RouterLinkProps } from 'react-router-dom';
import { LinkProps as MuiLinkProps } from '@material-ui/core/Link';
import { StyledLink } from './style';
import React from 'react';

type LinkPropsBase = RouterLinkProps & MuiLinkProps;

interface LinkProps extends Omit<LinkPropsBase, 'to'> {
  href?: string;
  to?: string;
}

const LinkRef = forwardRef<HTMLAnchorElement, RouterLinkProps>((props, ref) => (
  <RouterLink innerRef={ref} {...props} />
));

const Link = (props: LinkProps) => {
  const { href } = props;
  if (href) {
    return <StyledLink {...props} />;
  }
  return <StyledLink component={LinkRef} {...props} />;
};

Link.defaultProps = {
  href: null,
  to: '',
};

export default Link;
