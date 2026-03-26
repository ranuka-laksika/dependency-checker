# Dependency Checker with Claude AI

> **Purpose**: This repository is for getting approval for new dependencies through automated AI-powered analysis.

Automated dependency security and compliance analysis for Go projects using Claude AI. This system automatically analyzes Pull Requests that modify `go.mod` or `go.sum` files and provides comprehensive reports on:

- 🔒 Security vulnerabilities (CVEs)
- ⚖️ License compatibility (Apache 2.0 compliance)
- 📊 Dependency activity and freshness
- 🔗 Transitive dependency analysis

## Features

- **Automatic PR Analysis**: Automatically triggers when PRs modify dependency files
- **Instant Analysis**: Runs immediately when a PR is created or updated
- **Comprehensive Reports**: Checks each dependency for security, licensing, and maintenance status
- **Detailed Comments**: Adds structured analysis reports directly to PRs
- **Go-Specific**: Optimized for Go modules with `pkg.go.dev/vuln` integration

## How It Works

### Workflow Overview

```
PR Created/Updated (go.mod/go.sum)
           ↓
   Workflow Triggered Automatically
           ↓
   Claude Analyzes Dependencies
   (Security, License, Activity, Transitive)
           ↓
   Posts Detailed Report Comment on PR
           ↓
   Done!
```

### Simple & Direct

When a PR is opened or updated with changes to `go.mod` or `go.sum`:

- **File**: `.github/workflows/dependency_analysis.yml`
- **Trigger**: Automatically on PR opened/synchronized/reopened
- **Action**:
  - Runs Claude analysis immediately with specialized system prompt
  - Analyzes all dependency changes
  - Posts comprehensive report as PR comment
  - No labels, no queuing - just instant analysis

### System Prompt

Guides Claude's analysis behavior:

- **File**: `.github/claude/system_prompt.txt`
- **Purpose**: Instructs Claude on:
  - What to analyze (security, license, activity, transitive deps)
  - How to structure the report
  - Which tools to use (WebFetch, GitHub MCP tools, Go commands)
  - How to handle errors

## Setup Instructions

### Prerequisites

1. A GitHub repository with Go projects (`go.mod` files)
2. An Anthropic API key ([get one here](https://console.anthropic.com/))
3. GitHub repository admin access

### Step 1: Copy Files to Your Repository

Copy these files to your repository:

```
.github/
├── claude/
│   └── system_prompt.txt
└── workflows/
    └── dependency_analysis.yml
```

### Step 2: Configure GitHub Secrets

Go to your repository → Settings → Secrets and variables → Actions

Add the following secrets:

| Secret Name | Description | How to Get |
|------------|-------------|------------|
| `DEPENDENCY_CHECKER_ANTHROPIC_API_KEY` | Anthropic API key for Claude | [console.anthropic.com](https://console.anthropic.com/) |
| `DEPENDENCY_CHECKER_GITHUB_TOKEN` | GitHub PAT with repo access | See below |
| `DEPENDENCY_CHECKER_GIT_USER_NAME` | Git username | Your GitHub username |
| `DEPENDENCY_CHECKER_GIT_USER_EMAIL` | Git email | Your GitHub email |

#### Creating a GitHub Personal Access Token (PAT)

1. Go to GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Click "Generate new token (classic)"
3. Give it a descriptive name (e.g., "Dependency Checker Bot")
4. Select scopes:
   - ✅ `repo` (Full control of private repositories)
   - ✅ `workflow` (Update GitHub Action workflows)
5. Click "Generate token"
6. Copy the token and add it as `DEPENDENCY_CHECKER_GITHUB_TOKEN` secret

### Step 3: Enable GitHub Actions

1. Go to your repository → Settings → Actions → General
2. Under "Actions permissions", select:
   - ✅ "Allow all actions and reusable workflows"
3. Under "Workflow permissions", select:
   - ✅ "Read and write permissions"
   - ✅ "Allow GitHub Actions to create and approve pull requests"
4. Click "Save"

### Step 4: Test the Setup

1. Create a test branch:
   ```bash
   git checkout -b test-dependency-checker
   ```

2. Add a dependency to your `go.mod`:
   ```bash
   go get github.com/google/uuid@latest
   go mod tidy
   ```

3. Commit and push:
   ```bash
   git add go.mod go.sum
   git commit -m "test: add dependency for analysis"
   git push origin test-dependency-checker
   ```

4. Create a Pull Request

5. Watch the workflow run automatically:
   - Go to the PR's "Actions" tab or "Checks" section
   - The workflow will start immediately
   - Within minutes, a detailed analysis report will be posted as a comment

## Analysis Report Structure

Claude will post a comment with the following sections:

### Summary
- Total dependencies changed (added/updated/removed)
- Overall security risk level (🟢 Low / 🟡 Medium / 🔴 High)
- License compliance status
- Final recommendation (APPROVE / REVIEW / REJECT)

### Detailed Analysis (per dependency)

For each changed dependency, you'll get:

- **License Compatibility**: License type and Apache 2.0 compatibility status
- **Activity & Freshness**: Latest version, release date, repository activity
- **Security Posture**: CVEs, vulnerabilities from pkg.go.dev/vuln, govulncheck results
- **Transitive Dependencies**: Key transitive dependencies and their security status

### Final Recommendation
- Action items for the team
- Specific security or license concerns
- Upgrade recommendations if applicable

## Example Analysis Report

```markdown
# 🤖 Dependency Analysis Report

> **Analyzed by Claude AI** | PR #42 | Repository: myorg/myrepo

## 📋 Summary

- **Total Dependencies Changed**: 2
- **Added**: 1 | **Updated**: 1 | **Removed**: 0
- **Security Risk**: 🟢 Low
- **License Compliance**: ✅ All Compatible
- **Recommendation**: ✅ APPROVE

---

## 📦 Detailed Analysis

### github.com/google/uuid `v1.5.0` → `v1.6.0`

#### ✅ License Compatibility
- **License**: BSD-3-Clause
- **Apache 2.0 Compatible**: ✅ Yes
- **Notes**: Permissive license, safe for commercial use

#### 📊 Activity & Freshness
- **Latest Version**: v1.6.0 ✅ (this PR uses latest)
- **Release Date**: 2024-03-15 (2 weeks ago)
- **Repository Activity**: 🟢 Active (last commit: 2024-03-20)
- **Maintenance Status**: Well-maintained

#### 🔒 Security Posture
- **Vulnerabilities**: ✅ None found
- **govulncheck Status**: Clean
- **Recommendation**: ✅ Safe to use

#### 🔗 Transitive Dependencies
- No additional transitive dependencies

---

## 🎯 Final Recommendation

All dependency changes look good! No security concerns or license issues detected.

### ✅ Ready to Merge
- All licenses are Apache 2.0 compatible
- No known security vulnerabilities
- Using latest versions with active maintenance
```

## Customization

### Modify Report Format

Edit `.github/claude/system_prompt.txt` to change:
- Report structure and sections
- Analysis criteria and thresholds
- License compatibility rules
- Severity assessment logic
- Recommendation criteria

### Add Additional Checks

Modify the system prompt to add:
- Specific security policies for your organization
- Custom license requirements
- Minimum maintenance standards
- Specific dependency blocklists/allowlists

### Change Go Version

Edit `.github/workflows/dependency_analysis.yml`:

```yaml
- name: Setup Go
  uses: actions/setup-go@v5
  with:
    go-version: '1.23'  # Change to your version
```

### Add Support for Other Languages

To support npm (JavaScript), pip (Python), Maven (Java), etc.:

1. Update workflow trigger in `dependency_analysis.yml`:
   ```yaml
   paths:
     - 'go.mod'
     - 'go.sum'
     - 'package.json'      # npm
     - 'requirements.txt'  # pip
     - 'pom.xml'          # Maven
   ```

2. Modify `.github/claude/system_prompt.txt` with language-specific instructions

3. Add appropriate setup steps (Node.js, Python, Java) to the workflow

## Troubleshooting

### Workflow not triggering

- Check that the workflow file is on the default branch (main/master)
- Verify that the PR modifies `go.mod` or `go.sum`
- Check Actions tab → "All workflows" for any errors

### Claude not posting comments

- Verify `DEPENDENCY_CHECKER_ANTHROPIC_API_KEY` is set correctly
- Check that `DEPENDENCY_CHECKER_GITHUB_TOKEN` has required permissions
- Look for error messages in the workflow run logs (Actions tab)

### Analysis incomplete

- Check the workflow run logs in the Actions tab
- Common issues:
  - API rate limits (pkg.go.dev, GitHub)
  - Invalid or private dependencies
  - Network timeouts
- Claude will typically note what it couldn't analyze in the report

### Getting "Resource not accessible by integration" error

Update repository settings:
- Settings → Actions → General → Workflow permissions
- Enable "Read and write permissions"
- Enable "Allow GitHub Actions to create and approve pull requests"

### Analysis taking too long

The workflow typically completes in 2-5 minutes, but may take longer for PRs with many dependencies. If it consistently times out:
- Check if dependencies are accessible
- Verify network connectivity in workflow logs
- Consider splitting large dependency updates into smaller PRs

## Cost Estimation

Based on Claude Sonnet 4.5 pricing (March 2025):

- **Per PR analysis**: ~$0.50-2.00 (depends on number of dependencies)
- **Typical monthly cost**: $10-50 for active repositories (20-50 PRs/month)
- **Heavy usage**: ~$100-200/month for very active repos (100+ PRs/month)

Factors affecting cost:
- Number of dependencies changed
- Complexity of transitive dependency trees
- Number of external web fetches required

To reduce costs:
- Use Claude Haiku model for simpler analyses (edit model in workflow)
- Limit analysis to specific branches
- Only analyze production dependency changes (exclude dev dependencies)

## Security Considerations

- ✅ API keys are stored as GitHub Secrets (encrypted at rest)
- ✅ Bot only has read access to code, write access to comments
- ✅ Claude has no code modification capabilities (analysis only)
- ✅ All API calls are logged in Actions tab
- ✅ Workflow runs in isolated GitHub Actions environment
- ⚠️ Analysis reports are public if your repository is public

## Privacy

- Claude AI receives: PR details, go.mod content, and public package information
- Claude AI does NOT receive: Your source code, private repository data, or secrets
- All analysis uses publicly available data (pkg.go.dev, GitHub)

## Limitations

- Only analyzes Go modules (go.mod/go.sum) currently
- Depends on pkg.go.dev/vuln database (may have delays in CVE discovery)
- Cannot analyze private/internal dependencies without access
- Rate limited by external APIs (pkg.go.dev, GitHub)

## License

This project is designed for Apache 2.0 licensed repositories and checks dependencies for Apache 2.0 compatibility. Modify the system prompt to enforce different license requirements.

## Support

For issues or questions:
1. Check the [GitHub Actions logs](../../actions) for error details
2. Review Claude's analysis comments for specific issues
3. Verify all secrets and permissions are configured correctly
4. Open an issue in this repository

## Advanced Usage

### Manual Re-trigger

If you need to re-run the analysis:
1. Close and reopen the PR, OR
2. Push a new commit to the PR branch, OR
3. Go to Actions → Select the workflow run → "Re-run all jobs"

### Integration with Required Checks

Make the analysis a required status check:
1. Go to Settings → Branches → Branch protection rules
2. Select your protected branch (e.g., `main`)
3. Enable "Require status checks to pass before merging"
4. Search for and add "analyze-dependencies" to required checks

### Notifications

To get notified about high-risk dependencies:
- Watch for security-review-needed or license-review-needed labels (if you add label management to the workflow)
- Set up GitHub notifications for PR comments
- Create custom webhook integrations based on workflow status

## Credits

Inspired by the [WSO2 API Manager docs-fixing agent](https://github.com/wso2/docs-apim) pattern.

Built with [Claude Code](https://code.claude.com/) and [Claude AI](https://www.anthropic.com/claude).

---

**Questions?** Check the [Claude Code documentation](https://code.claude.com/docs) or [Anthropic's API docs](https://docs.anthropic.com/).
