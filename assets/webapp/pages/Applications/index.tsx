import React from 'react';
import { useParams } from 'react-router-dom';
import ApplicationsTable from '../../components/ApplicationsTable';
import useApplications from '../../hooks/useApplications';
import EnvironmentLayout from '../../layouts/Environment';


const ApplicationsPage = () => {
  const { envId } = useParams();
  const { status, error, data, isLoading } = useApplications(envId);
  const apps = data?.data;
  console.log(apps);
  return (
    <EnvironmentLayout>
      <ApplicationsTable apps={apps} />
    </EnvironmentLayout>
  )
};

export default ApplicationsPage;
