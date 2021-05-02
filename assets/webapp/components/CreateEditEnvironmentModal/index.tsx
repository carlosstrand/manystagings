import { ModalProps } from '@material-ui/core/Modal';
import React, { useEffect, useState } from 'react';
import useCreateEnvironment from '../../hooks/useCreateEnvironment';
import useEditEnvironment from '../../hooks/useEditEnvironment';
import useEnvironment from '../../hooks/useEnvironment';
import Button from '../../ui/Button';
import Input from '../../ui/Input';
import Modal from '../../ui/Modal';
import safeName from '../../utils/safeName';

interface CreateEditEnvironmentModalProps  {
  id: string | null;
  open: boolean;
  onClose: ModalProps["onClose"],
  refetch: () => void,
}

const CreateEditEnvironmentModal = (props: CreateEditEnvironmentModalProps) => {
  const { id, open, onClose } = props;
  const isEdit = !!id;
  const environment = useEnvironment(id);
  const editEnv = useEditEnvironment({});
  const createEnv = useCreateEnvironment({});
  const [name, setName] = useState('');
  const [namespace, setNamespace] = useState('');
  useEffect(() => {
    setName(environment.data?.name || '');
    setNamespace(environment.data?.namespace || '');
  }, [environment.data]);
  const onSubmit = () => {
    if (isEdit) {
      editEnv.mutateAsync({
        id,
        name,
        namespace,
      }).then(() => {
        environment.refetch();
        onClose({}, 'escapeKeyDown');
      })
    } else {
      createEnv.mutateAsync({
        name: name,
        namespace,
      }).then(() => {
        props.refetch();
        onClose({}, 'escapeKeyDown');
      })
    }
  };
  return (
    <Modal
      open={open}
      onClose={onClose}
      title={`${isEdit ? 'Edit' : 'Create'} Environment Variable`}
      loading={environment.isLoading}
      footer={(
        <div style={{ textAlign: 'right' }}>
          <Button onClick={() => onClose({}, 'escapeKeyDown')}>Cancel</Button>
          <Button
            color="primary"
            onClick={() => onSubmit()}
            disabled={name === "" || namespace === ""}
          >
            Save
          </Button>
        </div>
      )}
    >
      <div>
        <Input
          label="Name"
          value={name}
          onChange={(e) => {
            setName(e.target.value);
            setNamespace(safeName(e.target.value));
          }}
          fullWidth
        />
        <Input
          label="Namespace"
          value={namespace}
          onChange={(e) => {
            setNamespace(safeName(e.target.value));
          }}
          fullWidth
        />
      </div>
    </Modal>
  )
}

export default CreateEditEnvironmentModal;
