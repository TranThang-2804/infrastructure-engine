# Infrastructure Engine

## Overview

Infrastructure Engine is a self-service platform for managing infrastructure-as-code (IaC) pipelines. It is designed to work seamlessly with various IaC tools while following GitOps principles. This project aims to provide a flexible and integrated solution for teams already using IaC, offering better integration with existing systems compared to alternatives like Terraform controller with Kubevela/Crossplane.

## Architecture

The Infrastructure Engine is designed to integrate seamlessly with various Infrastructure-as-Code (IaC) tools, providing a unified platform for managing and orchestrating IaC pipelines. It leverages GitOps principles to ensure that infrastructure changes are version-controlled and automated.

- **Integration with IaC Tools**: The platform supports integration with popular IaC tools such as Terraform, Ansible, and CloudFormation, allowing users to manage their infrastructure using familiar tools.
- **GitOps Workflow**: Changes to infrastructure are managed through Git repositories, ensuring that all modifications are tracked, reviewed, and automated.
- **Self-Service Capability**: Users can provision and manage infrastructure resources through a self-service interface, reducing the need for manual intervention.
- **Scalability and Flexibility**: The architecture is designed to scale with the growing needs of the organization, supporting a wide range of infrastructure configurations and environments.

## Design Principles

The project is structured as a Go application with a layered architecture:

- **API Layer**: Exposes RESTful endpoints for managing blueprints, IaC templates, composite resources, and health checks.
- **Controller Layer**: Handles business logic and orchestrates interactions between the API and underlying services.
- **Use Case Layer**: Implements core business logic and workflows.
- **Repository Layer**: Manages data persistence and interactions with external systems (e.g., Git stores).
- **Infrastructure Layer**: Provides integrations with external services such as message queues (NATS) and Git providers.

## API Endpoints

- **Health Check**: `GET /health`
- **Blueprints**: `GET /blueprint`
- **IaC Templates**: `GET /iac-template`
- **Composite Resources**: `GET /composite`, `POST /composite`

## Setup and Installation

### Prerequisites

- Go 1.23.4 or higher
- [Devbox](https://get.jetify.com/devbox) for local development
- [Earthly](https://earthly.dev/) for CI/CD

### Local Development

1. **Install Development Tools**:
   ```bash
   task local:install
   ```

2. **Enter Development Shell**:
   ```bash
   task local:shell
   ```

3. **Bootstrap Local Environment**:
   ```bash
   task local:bootstrap
   ```

4. **Start Development Server**:
   ```bash
   task local:dev
   ```

5. **Clean Up Resources**:
   ```bash
   task local:clean
   ```

### Running Tests

```bash
task local:test
```

## CI/CD

The project uses Earthly for CI/CD. To run the CI pipeline locally, ensure you are in the development shell and execute:

```bash
task ci:execute
```

## License

This project is licensed under the terms of the LICENSE file included in the repository.

## Contact

For any questions or contributions, please contact Tommy Tran Duc Thang at [tranthang.dev@gmail.com](mailto:tranthang.dev@gmail.com).