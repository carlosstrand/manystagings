import React from 'react';
import MUIModal, { ModalProps as MUIModalProps } from '@material-ui/core/Modal';
import { Body, Dialogue, Header, Wrapper, Footer } from './style';


interface ModalProps  {
  children?: React.ReactNode;
  open: boolean;
  onClose: MUIModalProps["onClose"];
  loading?: boolean;
  title?: string;
  footer?: React.ReactNode;
}

const Modal = ({ children, open, onClose, loading, title, footer }: ModalProps) => {
  return (
    <MUIModal open={open} onClose={onClose} style={{display:'flex',alignItems:'center',justifyContent:'center'}}>
        <Wrapper>
          <Dialogue loading={loading} padding={0}>
            {title && (          
              <Header>
                <h4>{title}</h4>
              </Header>
            )}
            <Body>
              {children}
            </Body>
            {footer && (
              <Footer>
                {footer}
              </Footer>
            )}
          </Dialogue>
        </Wrapper>
    </MUIModal>
  )
};

export default Modal;
