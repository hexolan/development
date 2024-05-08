namespace Formulator.Core.Entities
{
    public class User : BaseEntity
    {
        public int Id { get; set; }

        public required string DisplayName { get; set; }

        public required string Email { get; set; }
    }
}
