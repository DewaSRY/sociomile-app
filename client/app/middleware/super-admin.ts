export default defineNuxtRouteMiddleware(async (to, from) => {
  const { isAuthenticated, loadUser, isSuperAdmin } = useAuth();
  
  if (!isAuthenticated.value) {
    await loadUser();
  }

  if (!isAuthenticated.value) {
    return navigateTo('/signin');
  }

  if (!isSuperAdmin()) {
    return navigateTo('/');
  }
});
