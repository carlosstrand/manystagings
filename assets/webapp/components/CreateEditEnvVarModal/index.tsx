import { ModalProps } from '@material-ui/core/Modal';
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import useCreateEnvVar from '../../hooks/useCreateEnvVar';
import useEditEnvVar from '../../hooks/useEditEnvVar';
import useEnvVar from '../../hooks/useEnvVar';
import Button from '../../ui/Button';
import Input from '../../ui/Input';
import Modal from '../../ui/Modal';

interface CreateEditEnvVarModalProps  {
  id: string | null;
  open: boolean;
  onClose: ModalProps["onClose"],
  refetch: () => void,
}

const CreateEditEnvVarModal = (props: CreateEditEnvVarModalProps) => {
  const { id, open, onClose } = props;
  const isEdit = !!id;
  const envVarQuery = useEnvVar(id);
  const editEnvVar = useEditEnvVar({});
  const createEnvVar = useCreateEnvVar({});
  const { appId } = useParams();
  const [envKey, setEnvKey] = useState('');
  const [envVal, setEnvVal] = useState('');
  useEffect(() => {
    setEnvKey(envVarQuery.data?.key || '');
    setEnvVal(envVarQuery.data?.value || '');
  }, [envVarQuery.data]);
  const onSubmit = () => {
    if (isEdit) {
      editEnvVar.mutateAsync({
        id,
        key: envKey,
        value: envVal,
      }).then(() => {
        props.refetch();
        onClose({}, 'escapeKeyDown');
      })
    } else {
      createEnvVar.mutateAsync({
        application_id: appId,
        key: envKey,
        value: envVal,
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
      loading={envVarQuery.isLoading}
      footer={(
        <div style={{ textAlign: 'right' }}>
          <Button onClick={() => onClose({}, 'escapeKeyDown')}>Cancel</Button>
          <Button color="primary" onClick={() => onSubmit()}>Save</Button>
        </div>
      )}
    >
      <div>
        <Input label="Key" value={envKey} onChange={(e) => setEnvKey(e.target.value)} fullWidth />
        <Input label="Value" value={envVal} onChange={(e) => setEnvVal(e.target.value)} fullWidth />
      </div>
    </Modal>
  )
}

export default CreateEditEnvVarModal;
