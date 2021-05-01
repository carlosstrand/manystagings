import React from 'react';
import SelectEnvironment from '../../components/SelectEnvironment';
import useEnvironments from '../../hooks/useEnvironments';
import CenteredBoxLayoyt from '../../layouts/CenteredBox';

const SelectEnvironmentPage = () => {
  const { status, error, data, isLoading, refetch } = useEnvironments();
  const environments = data?.data;
  return (
    <CenteredBoxLayoyt withLogo loading={isLoading}>
      {environments && <SelectEnvironment environments={environments} refetch={refetch} />}
    </CenteredBoxLayoyt>
  );
}

export default SelectEnvironmentPage;
