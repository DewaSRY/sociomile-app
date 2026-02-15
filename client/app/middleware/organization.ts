export default defineNuxtRouteMiddleware(async (to, from) => {
  const { isAuthenticated, loadUser, isOrganizationOwner, isOrganizationSales } = useAuth();
  
  if (!isAuthenticated.value) {
    await loadUser();
  }

  if (!isAuthenticated.value) {
    return navigateTo('/signin');
  }

  if (!isOrganizationOwner() && !isOrganizationSales()) {
    return navigateTo('/');
  }
});
