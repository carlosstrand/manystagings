import React from 'react';
import SelectEnvironment from '../../components/SelectEnvironment';
import useEnvironments from '../../hooks/useEnvironments';
import CenteredBoxLayoyt from '../../layouts/CenteredBox';

const SelectEnvironmentPage = () => {
  const { status, error, data, isLoading } = useEnvironments();
  const environments = data?.data;
  return (
    <CenteredBoxLayoyt withLogo loading={isLoading}>
      {environments && <SelectEnvironment environments={environments} />}
    </CenteredBoxLayoyt>
  );
}

export default SelectEnvironmentPage;
