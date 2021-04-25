import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';

interface AuthInput {
  username: string;
  password: string;
}

interface AuthPayload {
  token: {
    value: string;
    expiration: string;
  };
}

interface AuthResponse {
  auth: AuthPayload;
}

interface AuthRequest {
  input: AuthInput;
}

const useAuth = (opts: UseMutationOptions<AuthPayload, AxiosError, AuthInput>) =>
  useMutation<AuthPayload, AxiosError, AuthInput>(
    'auth',
    async (input: AuthInput) => client.post<AuthInput, AuthPayload>('/auth', {
      username: input.username,
      password: input.password,
    }),
    opts
  );

export default useAuth;
