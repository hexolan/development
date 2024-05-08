namespace Formulator.Core.Entities
{
    public class Form : BaseEntity
    {
        public int Id { get; set; }

        public int CreatorId { get; set; }
        public User? Creator { get; set; }

        public required string Title { get; set; }
        public string? Description { get; set; }

        public bool RequiresAuth { get; set; }

        // todo: review proper way of doing these
        public ICollection<FormQuestion> Questions { get; } = new List<FormQuestion>();
        public ICollection<FormSubmission> Submissions { get; } = new List<FormSubmission>();
    }
}
