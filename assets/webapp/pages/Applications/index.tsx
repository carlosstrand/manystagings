import React, { useState } from 'react';
import styled from 'styled-components';
import { useParams } from 'react-router-dom';
import ApplicationsTable from '../../components/ApplicationsTable';
import useApplications from '../../hooks/useApplications';
import EnvironmentLayout from '../../layouts/Environment';
import { StyledButton } from '../../ui/Button/style';
import { Add } from '@material-ui/icons';
import Button from '../../ui/Button';
import CreateEditApplicationModal from '../../components/CreateEditApplicationModal';

const Header = styled.div`
  margin-bottom: 16px;
  clear: both;
  h1 {
    display: inline-block;
  }
  ${StyledButton} {
    float: right;
    margin-top: 20px;
  }
`;

const ApplicationsPage = () => {
  const { envId } = useParams();
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const { status, error, data, isLoading, refetch } = useApplications(envId);
  const apps = data?.data;
  console.log(apps);
  return (
    <EnvironmentLayout>
      <CreateEditApplicationModal
        id={null}
        onClose={() => setCreateModalOpen(false)}
        open={createModalOpen}
        refetch={refetch}
      />
      <Header>
        <h1>
          Applications
        </h1>
        <Button onClick={() => setCreateModalOpen(true)}>
          <Add /> Create Application
        </Button>
      </Header>
      <ApplicationsTable apps={apps} />
    </EnvironmentLayout>
  )
};

export default ApplicationsPage;
