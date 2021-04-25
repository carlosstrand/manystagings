import React from 'react';
import styled from 'styled-components';
import { useParams } from 'react-router-dom';
import EnvVars from '../../components/EnvVars';
import useEnvVars from '../../hooks/useEnvVars';
import EnvironmentLayout from '../../layouts/Environment';

const ApplicationPage = () => {
  return (
    <EnvironmentLayout>
      <EnvVars />
    </EnvironmentLayout>
  )
};

export default ApplicationPage;
