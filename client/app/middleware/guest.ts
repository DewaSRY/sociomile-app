export default defineNuxtRouteMiddleware(async (to, from) => {
  const { isAuthenticated, loadUser } = useAuth();
  
  // Load user if token exists
  if (!isAuthenticated.value) {
    await loadUser();
  }

  // Redirect authenticated users to appropriate dashboard
  if (isAuthenticated.value) {
    const { isSuperAdmin, isOrganizationOwner, isOrganizationSales, isGuest } = useAuth();
    
    if (isSuperAdmin()) {
      return navigateTo('/hub/dashboard');
    } else if (isOrganizationOwner() || isOrganizationSales()) {
      return navigateTo('/organization/dashboard');
    } else if (isGuest()) {
      return navigateTo('/guest/dashboard');
    }
  }
});
