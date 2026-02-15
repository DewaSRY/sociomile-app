<template>
  <div>
    <UPageHero
      title="Sociomile - Customer Communication Hub"
      description="Manage all your customer conversations in one place. Track messages from WhatsApp, Instagram, and other channels. Assign conversations to your sales team and turn unresolved issues into trackable tickets."
      :links="[
        {
          label: 'Get started',
          to: '/auth/signin',
          trailingIcon: 'i-lucide-arrow-right',
          size: 'xl',
        },
        {
          label: 'Sign in',
          to: '/auth/signup',
          icon: 'i-lucide-log-in',
          size: 'xl',
          color: 'neutral',
          variant: 'subtle',
        },
      ]"
    />

    <UPageSection
      id="features"
      title="Everything you need to manage customer conversations"
      description="A powerful multi-tenant platform designed to help business owners track, manage, and respond to customer messages efficiently."
      :features="[
        {
          icon: 'i-lucide-message-circle',
          title: 'Centralized Conversations',
          description:
            'Capture and organize all customer messages from multiple channels (WhatsApp, Instagram, etc.) in a unified dashboard.',
        },
        {
          icon: 'i-lucide-users',
          title: 'Team Assignment',
          description:
            'Business owners can assign conversations to specific sales staff members, ensuring every customer gets proper attention and follow-up.',
        },
        {
          icon: 'i-lucide-ticket',
          title: 'Ticket Management',
          description:
            'Convert unresolved customer issues into trackable tickets. Monitor progress and ensure no customer complaint falls through the cracks.',
        },
        {
          icon: 'i-lucide-building-2',
          title: 'Multi-tenant Architecture',
          description:
            'Each organization operates independently with its own dashboard, staff, and customer conversations in a secure, isolated environment.',
        },
        {
          icon: 'i-lucide-layout-dashboard',
          title: 'Role-based Dashboards',
          description:
            'Different views for super admins, organization owners, sales staff, and guests. Everyone sees exactly what they need.',
        },
        {
          icon: 'i-lucide-shield-check',
          title: 'Secure & Scalable',
          description:
            'Built with Go and MySQL for enterprise-grade performance, security, and reliability as your business grows.',
        },
      ]"
    />

    <UPageSection>
      <UPageCTA
        title="Ready to streamline your customer communication?"
        description="Join businesses using Sociomile to manage customer conversations efficiently. Start tracking, assigning, and resolving customer issues today."
        variant="subtle"
        :links="[
          {
            label: 'Create an account',
            to: '/auth/register',
            trailingIcon: 'i-lucide-arrow-right',
            color: 'neutral',
          },
          {
            label: 'Learn more',
            to: '/about',
            icon: 'i-lucide-info',
            color: 'neutral',
            variant: 'outline',
          },
        ]"
      />
    </UPageSection>
  </div>
</template>
