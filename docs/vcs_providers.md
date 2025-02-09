# VCS Providers

To connect workspaces and modules to git repositories containing Terraform configurations, you need to provide OTF with access to your VCS provider.

Firstly, create a provider for your organization. On your organization's main menu, select **VCS providers**.

![organization main menu](images/organization_main_menu.png){.screenshot}

You'll be presented with a choice of providers to create. The choice is restricted to those for which you have enabled [SSO](#authentication). For instance, if you have enabled Github SSO then you can create a Github VCS provider.

![vcs providers list](images/vcs_providers_list.png){.screenshot}

Select the provider you would like to create. You will then be prompted to enter a personal access token. Instructions for generating the token are included on the page. The token permits OTF to access your git repository and retrieve terraform configuration. Once you've generated and inserted the token into the field you also need to give the provider a name that describes it.

![new github vcs provider form](images/new_github_vcs_provider_form.png){.screenshot}

!!! note
    Be sure to restrict the permissions on the token according to the instructions.

Create the provider and it'll appear on the list of providers. You can now proceed to connecting workspaces and publishing modules.

### Connecting a workspace

Once you have a provider you can connect a workspace to a git repository for that provider.

Select a workspace. Go to **settings**.

![workspace main page](images/workspace_main_page.png){.screenshot}

Click **Connect to VCS**.

![workspace settings](images/workspace_settings.png){.screenshot}

Select the provider.

![workspace vcs providers list](images/workspace_vcs_providers_list.png){.screenshot}

You'll then be presented with a list of repositories. Select the repository containing the terraform configuration you want to use in your workspace. If you cannot see your repository you can enter its name.

![workspace vcs repo list](images/workspace_vcs_repo_list.png){.screenshot}

Once connected you can start a run via the web UI. On the workspace page select the **start run** drop-down box and select an option to either start a plan or both a plan and an apply.

![connected workspace main page](images/connected_workspace_main_page.png){.screenshot}

That will start a run, retrieving the configuration from the repository, and you will see the progress of its plan and apply.

![run page started](images/run_page_started.png){.screenshot}
