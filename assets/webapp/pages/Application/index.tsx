import React from 'react';
import { useParams } from 'react-router-dom';
import EnvVars from '../../components/EnvVars';
import useEnvVars from '../../hooks/useEnvVars';
import EnvironmentLayout from '../../layouts/Environment';

const ApplicationPage = () => {
  const { appId } = useParams();
  const envQuery = useEnvVars(appId);
  console.log(envQuery);
  return (
    <EnvironmentLayout>
      <EnvVars envVars={envQuery?.data?.data || []} />
    </EnvironmentLayout>
  )
};

export default ApplicationPage;
