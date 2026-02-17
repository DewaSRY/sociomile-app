import { API_AUTH_PROFILE } from "~~/shared/constants/api-path";

export const useProfile = () => {
  const profile = useState<UserProfileData | null>("profile", () => null);
  const pending = useState<boolean>("profile-pending", () => false);

  const fetchProfile = async () => {
    if (pending.value) return;

    pending.value = true;

    try {
      const data = await $fetch<UserProfileData>(API_AUTH_PROFILE);
      profile.value = data;
    } catch (err) {
      profile.value = null;
    } finally {
      pending.value = false;
    }

    return profile.value;
  };

  const reset = () => {
    profile.value = null;
  };

  return {
    profile,
    fetchProfile,
    pending,
    reset,
  };
};
