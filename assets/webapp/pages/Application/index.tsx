import { Delete, Settings } from '@material-ui/icons';
import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import styled from 'styled-components';
import ApplicationPropertiesTable from '../../components/ApplicationPropertiesTable';
import CreateEditApplicationModal from '../../components/CreateEditApplicationModal';
import EnvVars from '../../components/EnvVars';
import useApplication from '../../hooks/useApplication';
import useDeleteApplication from '../../hooks/useDeleteApplication';
import EnvironmentLayout from '../../layouts/Environment';
import Box from '../../ui/Box';
import Button from '../../ui/Button';
import { StyledButton } from '../../ui/Button/style';

const Header = styled.div`
  margin-bottom: 16px;
  clear: both;
  h1 {
    display: inline-block;
  }
  ${StyledButton} {
    float: right;
    margin-top: 20px;
    margin-left: 8px;
  }
`;

const AppSection = styled(Box)`
  margin-bottom: 48px;
`;

const ApplicationPage = () => {
  const { appId } = useParams();
  const [showSettings, setShowSettings] = useState(false);
  const { data, isLoading, refetch } = useApplication(appId);
  const deleteApp = useDeleteApplication({});
  const propertiesRows = !data ? [] : [
    {
      key: "Docker Image",
      value: data.docker_image_name,
    },
    {
      key: "Docker Tag",
      value: data.docker_image_tag,
    },
    {
      key: "Port",
      value: data.port.toString(),
    },
    {
      key: "Container Port",
      value: data.container_port.toString(),
    },
    {
      key: "Public URL",
      value: data.public_url.toString(),
    }
  ];
  return (
    <EnvironmentLayout>
      <CreateEditApplicationModal
        id={appId}
        onClose={() => setShowSettings(false)}
        open={showSettings}
        refetch={refetch}
      />
      <AppSection loading={isLoading}>
        <Header>
          <h1>{data?.name}</h1>
          <Button color="secondary" onClick={(e) => {
            e.preventDefault();
            const confirmed = confirm(`Are you sure you want to delete the application ${data?.name}?`);
            if (confirmed) {
              deleteApp.mutateAsync({
                id: appId,
              }).then(() => {
                refetch();
              })
            }
          }}>
            <Delete /> Delete
          </Button>
          <Button onClick={() => setShowSettings(true)}>
            <Settings /> Settings
          </Button>
        </Header>
        <ApplicationPropertiesTable rows={propertiesRows} />
      </AppSection>
      <AppSection loading={isLoading}>
        <EnvVars />
      </AppSection>
    </EnvironmentLayout>
  )
};

export default ApplicationPage;
