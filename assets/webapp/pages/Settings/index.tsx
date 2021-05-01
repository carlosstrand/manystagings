import React, { useState } from 'react';
import styled from 'styled-components';
import { useParams, useHistory } from 'react-router-dom';
import ApplicationsTable from '../../components/ApplicationsTable';
import EnvironmentLayout from '../../layouts/Environment';
import Button from '../../ui/Button';
import useDeleteEnvironment from '../../hooks/useDeleteEnvironment';
import { Settings } from '@material-ui/icons';
import CreateEditEnvironmentModal from '../../components/CreateEditEnvironmentModal';

const Header = styled.div`
  margin-bottom: 16px;
  clear: both;
  h1 {
    display: inline-block;
  }
`;

const SettingsPage = () => {
  const [showUpdateEnvironmentModal, setShowUpdateEnvironmentModal] = useState(false);
  const { envId } = useParams();
  const history = useHistory();
  const deleteEnvironment = useDeleteEnvironment({});
  return (
    <EnvironmentLayout>
      <CreateEditEnvironmentModal
        id={envId}
        open={showUpdateEnvironmentModal}
        onClose={() => setShowUpdateEnvironmentModal(false)}
        refetch={() => {}}
      />
      <Header>
        <h1>
          Settings
        </h1>
      </Header>

      <h2>
        General
      </h2>
      <Button onClick={(e) => setShowUpdateEnvironmentModal(true)}>
        <Settings /> Settings
      </Button>

      <h2>
        Danger Zone
      </h2>
      <Button color="secondary" onClick={(e) => {
        e.preventDefault();
        const confirmed = confirm(`Are you sure you want to delete the environment?`);
        if (confirmed) {
          deleteEnvironment.mutateAsync({
            id: envId,
          }).then(() => {
            history.push('/select-environment');
          })
        }
      }}>
        Delete Environment
      </Button>
    </EnvironmentLayout>
  )
};

export default SettingsPage;
