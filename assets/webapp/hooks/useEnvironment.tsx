import { useQuery } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import Environment from '../types/environment';

const useEnvironment = (id?: string) =>
  useQuery<Environment, AxiosError>(
    ['fetchEnvironment', id],
    async () => (await client.get(`/api/environments/${id}`)).data,
    {
      enabled: !!id,
    }
  );

export default useEnvironment;
