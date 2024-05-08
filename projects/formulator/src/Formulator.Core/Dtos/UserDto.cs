namespace Formulator.Application.DTOs
{
    public class UserDTO : BaseDto
    {
        public int Id { get; set; }
        public required string Email { get; set; }
        public required string DisplayName { get; set; }
    }
}
