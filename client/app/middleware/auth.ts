export default defineNuxtRouteMiddleware(async (to, from) => {
  const { isAuthenticated, loadUser } = useAuth();
  
  // Load user if token exists but user is not loaded
  if (!isAuthenticated.value) {
    await loadUser();
  }

  // Redirect to signin if not authenticated
  if (!isAuthenticated.value) {
    return navigateTo('/signin');
  }
});
