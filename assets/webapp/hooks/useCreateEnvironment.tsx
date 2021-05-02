import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import Environment from '../types/environment';

interface CreateEnvironmentInput {
  name: string;
  namespace: string;
}

const useCreateEnvironment = (opts: UseMutationOptions<Environment, AxiosError, CreateEnvironmentInput>) =>
  useMutation<Environment, AxiosError, CreateEnvironmentInput>(
    'createEnvironment',
    async (input: CreateEnvironmentInput) => client.post<CreateEnvironmentInput, Environment>('/api/environments', input),
    opts
  );

export default useCreateEnvironment;
