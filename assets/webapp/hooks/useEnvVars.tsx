import { useQuery } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import ApplicationEnvVar from '../types/application_env_var';

interface ApplicationEnvVars {
  data: ApplicationEnvVar[];
  count: number;
}

const useEnvVars = (appId: string) =>
  useQuery<ApplicationEnvVars, AxiosError>(
    'fetchEnvVars',
    async () => (await client.get('/api/application_env_vars', {
      params: {
        filter: JSON.stringify({
          where: {
            application_id: {
              eq: appId,
            },
          },
        }),
      },
    })).data,
  );

export default useEnvVars;
