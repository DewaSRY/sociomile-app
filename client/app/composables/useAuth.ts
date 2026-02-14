import type { AuthResponse, UserData } from '~/types';

interface AuthState {
  user: UserData | null;
  token: string | null;
  isAuthenticated: boolean;
}

export const useAuth = () => {
  const authState = useState<AuthState>('auth', () => ({
    user: null,
    token: null,
    isAuthenticated: false,
  }));

  const token = useCookie<string | null>("auth_token", {
    maxAge: 60 * 60 * 24 * 7, // 7 days
    sameSite: "lax",
    secure: false, // process.env.NODE_ENV === 'production
  });

  const setAuth = (authData: AuthResponse) => {
    authState.value.user = authData.user;
    authState.value.token = authData.token;
    authState.value.isAuthenticated = true;
    token.value = authData.token;
  };

  const clearAuth = () => {
    authState.value.user = null;
    authState.value.token = null;
    authState.value.isAuthenticated = false;
    token.value = null;
  };

  const loadUser = async () => {
    if (token.value && !authState.value.user) {
      try {
        const user = await $fetch<UserData>('/api/auth/me', {
          headers: {
            Authorization: `Bearer ${token.value}`,
          },
        });
        authState.value.user = user;
        authState.value.token = token.value;
        authState.value.isAuthenticated = true;
      } catch (error) {
        clearAuth();
      }
    }
  };

  const getUserRole = () => {
    return authState.value.user?.role?.name || null;
  };

  const isRole = (role: string) => {
    return getUserRole() === role;
  };

  const isSuperAdmin = () => isRole('super_admin');
  const isOrganizationOwner = () => isRole('organization_owner');
  const isOrganizationSales = () => isRole('organization_sales');
  const isGuest = () => isRole('guest');

  return {
    user: computed(() => authState.value.user),
    token: computed(() => authState.value.token || token.value),
    isAuthenticated: computed(() => authState.value.isAuthenticated),
    setAuth,
    clearAuth,
    loadUser,
    getUserRole,
    isRole,
    isSuperAdmin,
    isOrganizationOwner,
    isOrganizationSales,
    isGuest,
  };
};
