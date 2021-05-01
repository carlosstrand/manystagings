import { ModalProps } from '@material-ui/core/Modal';
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import useApplication from '../../hooks/useApplication';
import useCreateApplication from '../../hooks/useCreateApplication';
import useEditApplication from '../../hooks/useEditApplication';
import Application from '../../types/application';
import Button from '../../ui/Button';
import Input from '../../ui/Input';
import Modal from '../../ui/Modal';

interface CreateEditApplicationModalProps  {
  id: string | null;
  open: boolean;
  onClose: ModalProps["onClose"],
  refetch: () => void,
}

const CreateEditApplicationModal = (props: CreateEditApplicationModalProps) => {
  const { id, open, onClose } = props;
  const isEdit = !!id;
  const appQuery = useApplication(id);
  const editApp = useEditApplication({});
  const createApp = useCreateApplication({});
  const { envId, appId } = useParams();
  const [form, setForm] = useState<any>({
    name: '',
    docker_image_name: '',
    docker_image_tag: '',
    port: '',
    container_port: '',
  });
  const formToInput = (form): Application => ({
    environment_id: envId,
    name: form.name,
    docker_image_name: form.docker_image_name,
    docker_image_tag: form.docker_image_tag,
    port: parseInt(form.port, 10),
    container_port: parseInt(form.container_port, 10),
  })
  useEffect(() => {
    setForm(appQuery.data || {});
  }, [appQuery.data]);
  const onSubmit = () => {
    if (isEdit) {
      const data = formToInput(form);
      data.id = appId;
      editApp.mutateAsync(data).then(() => {
        props.refetch();
        onClose({}, 'escapeKeyDown');
      })
    } else {
      createApp.mutateAsync(formToInput(form)).then(() => {
        props.refetch();
        onClose({}, 'escapeKeyDown');
      })
    }
  };
  const getFormProps = (field: any) => ({
    value: form[field],
    onChange: (e) => setForm(form => ({ ...form, [field]: e.target.value })),
  })
  return (
    <Modal
      open={open}
      onClose={onClose}
      title={`${isEdit ? 'Edit' : 'Create'} Application`}
      loading={appQuery.isLoading}
      footer={(
        <div style={{ textAlign: 'right' }}>
          <Button onClick={() => onClose({}, 'escapeKeyDown')}>Cancel</Button>
          <Button color="primary" onClick={() => onSubmit()}>Save</Button>
        </div>
      )}
    >
      <div>
        <Input label="Name" fullWidth {...getFormProps('name')} />
        <Input label="Docker Image Name" fullWidth {...getFormProps('docker_image_name')} />
        <Input label="Docker Image Tag" fullWidth {...getFormProps('docker_image_tag')} />
        <Input label="Port" fullWidth {...getFormProps('port')} />
        <Input label="Container Port" fullWidth {...getFormProps('container_port')} />
      </div>
    </Modal>
  )
}

export default CreateEditApplicationModal;
