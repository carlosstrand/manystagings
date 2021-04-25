import { useQuery } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import EnvVar from '../types/application_env_var';

const useEnvVar = (id?: string) =>
  useQuery<EnvVar, AxiosError>(
    'fetchEnvVar',
    async () => (await client.get(`/api/application_env_vars/${id}`)).data,
    {
      enabled: !!id,
    }
  );

export default useEnvVar;
