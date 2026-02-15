export default defineNuxtRouteMiddleware(async (to, from) => {
  const { isAuthenticated, loadUser, isGuest } = useAuth();
  
  if (!isAuthenticated.value) {
    await loadUser();
  }

  if (!isAuthenticated.value) {
    return navigateTo('/signin');
  }

  if (!isGuest()) {
    return navigateTo('/');
  }
});
