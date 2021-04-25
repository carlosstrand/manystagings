export default interface Application {
  id: string;
  name: string;
  docker_image_name: string;
  docker_image_tag: string;
  shell_command: string;
  port: number;
  container_port: number;
}
